package repository

import (
	"flyAPI/internal/models"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type RouteRepo struct {
	db *sqlx.DB
}

func NewRouteRepo(db *sqlx.DB) *RouteRepo {
	return &RouteRepo{
		db: db,
	}
}

type RouteParams struct {
	DepartureAirport string
	Origin           string
	FareCondition    string
	LimitLength      int
	Date             time.Time
}

func (s *RouteRepo) GetRoutesFromAirport(aiport, startDate, endDate string) ([]models.Flight, error) {
	var flights []models.Flight
	query := fmt.Sprintf(`SELECT DISTINCT * from %s v WHERE v.departure_airport = $1 AND v.status = 'Scheduled' AND scheduled_departure > $2`, flightsTable)

	if err := s.db.Select(&flights, query, aiport, startDate); err != nil {
		return []models.Flight{}, err
	}

	return flights, nil
}
