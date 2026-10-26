package driver

import (
	"database/sql"
	"errors"
	"time"

	"github.com/doors-dev/doors"
)

func newSessionsDB(db *sql.DB) sessionsDB {
	initQuery := `
		CREATE TABLE IF NOT EXISTS sessions (
			token TEXT PRIMARY KEY,
			login TEXT NOT NULL,
			expire DATETIME NOT NULL
		);
	`
	if _, err := db.Exec(initQuery); err != nil {
		panic("Failed to create sessions table: " + err.Error())
	}
	s := sessionsDB{
		db: db,
	}
	go s.cleanup()
	return s
}

type Session struct {
	Token  string    `json:"token"`
	Login  string    `json:"login"`
	Expire time.Time `json:"expire"`
}

func (s Session) IsValid() bool {
	return s.Token != ""
}

type sessionsDB struct {
	db *sql.DB
}

func (d sessionsDB) cleanup() {
	for {
		<-time.After(10 * time.Minute)
		_, err := d.db.Exec("DELETE FROM sessions WHERE expire <= ?", time.Now())
		if err != nil {
			panic("Failed to cleanup expired sessions: " + err.Error())
		}
	}
}

func (d sessionsDB) Add(login string, dur time.Duration) Session {
	token := doors.IDRand()
	expire := time.Now().Add(dur)

	_, err := d.db.Exec(
		"INSERT INTO sessions (token, login, expire) VALUES (?, ?, ?)",
		token, login, expire,
	)
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}

	return Session{
		Token:  token,
		Login:  login,
		Expire: expire,
	}
}

func (d sessionsDB) Get(token string) Session {
	var session Session
	err := d.db.QueryRow(
		"SELECT token, login, expire FROM sessions WHERE token = ? AND expire > ?",
		token, time.Now(),
	).Scan(&session.Token, &session.Login, &session.Expire)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Session{}
		}
		panic("Failed to get session: " + err.Error())
	}

	return session
}

func (d sessionsDB) Remove(token string) bool {
	result, err := d.db.Exec("DELETE FROM sessions WHERE token = ?", token)
	if err != nil {
		panic("Failed to remove session: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic("Failed to get rows affected: " + err.Error())
	}

	return rowsAffected > 0
}
