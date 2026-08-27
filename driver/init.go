package driver

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

var Locations locationsDB
var Weather weatherAPI
var Sessions sessionsDB

func init() {
	locations, err := sql.Open("sqlite", "./sqlite/sqlite-world.sqlite3")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	if err := locations.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	Locations = locationsDB{db: locations}

	Weather = weatherAPI{
		endpoint: "https://api.open-meteo.com/v1/forecast",
		timeout:  10 * time.Second,
	}

	sessions, err := sql.Open("sqlite", "./sqlite/sessions.sqlite3")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	Sessions = newSessionsDB(sessions)
}
