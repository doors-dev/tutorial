package driver

import (
	"database/sql"
	"errors"
)

type Place struct {
	Id   int
	Name string
}

func (p Place) IsValid() bool {
	return p.Name != ""
}

type City struct {
	Name    string
	Country Place
	Id      int
	Lat     float64
	Long    float64
}

func (c City) IsValid() bool {
	return c.Name != ""
}

type locationsDB struct {
	db *sql.DB
}

func (d locationsDB) CountriesGet(id int) (Place, error) {
	var p Place
	query := `
		SELECT id, name
		FROM countries
		WHERE id = ?
		LIMIT 1
	`
	row := d.db.QueryRow(query, id)
	if err := row.Scan(&p.Id, &p.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Place{}, nil
		}
		return Place{}, err
	}
	return p, nil
}

func (d locationsDB) CountriesSearch(term string) ([]Place, error) {
	query := `
		SELECT id, name
		FROM countries
		WHERE LOWER(name) LIKE LOWER(?)
		LIMIT 7
	`
	rows, err := d.db.Query(query, term+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Place
	for rows.Next() {
		var c Place
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		results = append(results, c)
	}
	return results, nil
}

func (d locationsDB) CitiesGet(city int) (City, error) {
	var c City
	row := d.db.QueryRow(`
		SELECT id, name, latitude, longitude, country_id
		FROM cities
		WHERE id = ?
		LIMIT 1
	`, city)

	var countryId int
	var err error
	if err = row.Scan(&c.Id, &c.Name, &c.Lat, &c.Long, &countryId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return City{}, nil
		}
		return City{}, err
	}
	c.Country, err = d.CountriesGet(countryId)
	if err != nil {
		return City{}, err
	}
	return c, nil
}

func (d locationsDB) CitiesSearch(country int, term string) ([]Place, error) {
	query := `
		SELECT id, name
		FROM cities
		WHERE country_id = ?
		AND LOWER(name) LIKE LOWER(?)
		LIMIT 7
	`
	rows, err := d.db.Query(query, country, term+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Place
	for rows.Next() {
		var c Place
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		results = append(results, c)
	}
	return results, nil
}
