package repository

import (
	"flyAPI/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AirRepo struct {
	db *sqlx.DB
}

func NewAirRepo(db *sqlx.DB) *AirRepo {
	return &AirRepo{
		db: db,
	}
}

type Cities struct {
	SrcCities  []string
	DestCities []string
}

func (a *AirRepo) GetAllSrcAndDestCities() (Cities, error) {
	var cities Cities

	query1 := fmt.Sprintf("SELECT DISTINCT(departure_city) from %s", routeTable)

	if err := a.db.Select(&cities.SrcCities, query1); err != nil {
		return Cities{}, err
	}

	query2 := fmt.Sprintf("SELECT DISTINCT(arrival_city) from %s", routeTable)

	if err := a.db.Select(&cities.DestCities, query2); err != nil {
		return Cities{}, err
	}
	return cities, nil
}

type Airports struct {
	SrcAirports  []models.Airport
	DestAirports []models.Airport
}

func (a *AirRepo) GetAllSrcAndDestAirports() (Airports, error) {
	var airports Airports

	query1 := fmt.Sprintf(`
		SELECT DISTINCT(airport_code), airport_name, 
			city, coordinates, timezone
		FROM %s r left join %s a
		ON r.departure_airport = a.airport_code`, routeTable, airportDataTable)

	if err := a.db.Select(&airports.SrcAirports, query1); err != nil {
		return Airports{}, err
	}

	query2 := fmt.Sprintf(`
		SELECT DISTINCT(airport_code), airport_name, 
			city, coordinates, timezone
		FROM %s r left join %s a
		ON r.arrival_airport = a.airport_code`, routeTable, airportDataTable)

	if err := a.db.Select(&airports.DestAirports, query2); err != nil {
		return Airports{}, err
	}

	return airports, nil
}

func (a *AirRepo) GetAirportsInCity(cityName string) ([]models.Airport, error) {
	var airports []models.Airport

	query := fmt.Sprintf(`
		SELECT airport_code, airport_name, 
			city, coordinates, timezone
		FROM %s a WHERE position($1 in a.city::Text) > 0`, airportDataTable)

	if err := a.db.Select(&airports, query, cityName); err != nil {
		return nil, err
	}

	return airports, nil
}

func (a *AirRepo) GetAirportByName(airportName string) (models.Airport, error) {
	var airport models.Airport

	query := fmt.Sprintf(`
		SELECT airport_code, airport_name, 
			city, coordinates, timezone
		FROM %s a WHERE position($1 in a.airport_name::Text) > 0`, airportDataTable)

	if err := a.db.Get(&airport, query, airportName); err != nil {
		return models.Airport{}, err
	}

	return airport, nil
}
