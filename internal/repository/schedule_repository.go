package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ScheduleRepo struct {
	db *sqlx.DB
}

func NewScheduleRepo(db *sqlx.DB) *ScheduleRepo {
	return &ScheduleRepo{
		db: db,
	}
}

type InboundSchedule struct {
	DayOfWeek int     `db:"day_of_week" json:"day_of_week"`
	Time      string  `db:"time_of_arrival" json:"time_of_arrival"`
	FlightNo  []uint8 `db:"flight_no" json:"flight_no"`
	Origin    string  `db:"origin" json:"origin"`
}

func (s *ScheduleRepo) GetInboundScheduleForAirport(airport, time string) ([]InboundSchedule, error) {
	var schedules []InboundSchedule

	query := fmt.Sprintf(`
	SELECT
		EXTRACT(DOW FROM f.scheduled_arrival) AS day_of_week,
		f.scheduled_arrival::time AS time_of_arrival,
		f.flight_no,
		f.departure_airport AS origin
	FROM
		%s f
	WHERE
		f.arrival_airport = $1
		AND f.scheduled_arrival >= $2
	ORDER BY
		f.scheduled_arrival;
	`, flightsTable)

	if err := s.db.Select(&schedules, query, airport, time); err != nil {
		return nil, err
	}

	return schedules, nil
}

type OutboundSchedule struct {
	DayOfWeek int     `db:"day_of_week" json:"day_of_week"`
	Time      string  `db:"time_of_departure" json:"time_of_departure"`
	FlightNo  []uint8 `db:"flight_no" json:"flight_no"`
	Origin    string  `db:"destination" json:"destination"`
}

func (s *ScheduleRepo) GetOutboundScheduleForAirport(airport string, time string) ([]OutboundSchedule, error) {
	var schedules []OutboundSchedule

	query := fmt.Sprintf(`
	 SELECT
            EXTRACT(DOW FROM f.scheduled_departure) AS day_of_week,
            f.scheduled_departure::time AS time_of_departure,
            f.flight_no,
            f.arrival_airport AS destination
        FROM
            %s f
        WHERE
            f.departure_airport = $1
            AND f.scheduled_departure >= $2
        ORDER BY
            f.scheduled_departure;
	`, flightsTable)

	if err := s.db.Select(&schedules, query, airport, time); err != nil {
		return nil, err
	}

	return schedules, nil
}
