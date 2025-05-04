package repository

import (
	"flyAPI/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SeatRepo struct {
	db *sqlx.DB
}

func NewSeatRepo(db *sqlx.DB) *SeatRepo {
	return &SeatRepo{
		db: db,
	}
}

func (p *SeatRepo) GetSeatsByAircraftCode(aircraftCode string) (models.AircraftSeatsInfo, error) {
	var info models.AircraftSeatsInfo

	query := fmt.Sprintf(`SELECT s.aircraft_code, COUNT(s) as seat_amount, 
            SUM(CASE WHEN s.fare_conditions = 'Economy' THEN 1 ELSE 0 END) as econom_amount,
            SUM(CASE WHEN s.fare_conditions = 'Business' THEN 1 ELSE 0 END) as business_amount, 
            SUM(CASE WHEN s.fare_conditions = 'Comfort' THEN 1 ELSE 0 END)) as comfort_amount
            FROM %s s WHERE s.aircraft_code = $1`, seatsTable)

	err := p.db.Get(&info, query, aircraftCode)
	return info, err
}

func (p *SeatRepo) FindSeatAmountByAircraftCodeAndFareCondition(aircraftCode, fareConditions string) (models.AircraftSeatsByFareCondition, error) {
	var info models.AircraftSeatsByFareCondition

	query := fmt.Sprintf(`SELECT s.aircraft_code, SUM(CASE WHEN s.fare_conditions = $1 THEN 1 ELSE 0 END) as amount FROM %s s WHERE s.aircraft_code = $2`, seatsTable)

	err := p.db.Get(&info, query, fareConditions, aircraftCode)
	return info, err
}

func (p *SeatRepo) FindSeatsByAircraftCodeAndFareCondition(aircraftCode, fareConditions string) ([]models.Seat, error) {
	var info []models.Seat

	query := fmt.Sprintf(`SELECT * FROM %s s WHERE s.aircraft_code = $1 AND s.fare_condition = $2`, seatsTable)

	err := p.db.Select(&info, query, fareConditions, aircraftCode, fareConditions)
	return info, err
}
