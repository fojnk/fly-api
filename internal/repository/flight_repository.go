package repository

import (
	"flyAPI/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type FlightRepo struct {
	db *sqlx.DB
}

func NewFlightRepo(db *sqlx.DB) *FlightRepo {
	return &FlightRepo{
		db: db,
	}
}

func (f *FlightRepo) GetFlightByFlightId(flightId int) (models.Flight, error) {
	var flight models.Flight
	query := fmt.Sprintf(`SELECT * from %s f WHERE f.flight_id = $1`, flightsTable)

	err := f.db.Get(&flight, query, flightId)
	return flight, err
}
