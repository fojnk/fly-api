package models

import "database/sql"

type Flight struct {
	FlightId           int64          `db:"flight_id"`
	FlightNo           string         `db:"flight_no"`
	ScheduledDeparture string         `db:"scheduled_departure"`
	ScheduledArrival   string         `db:"scheduled_arrival"`
	DepartureAirport   string         `db:"departure_airport"`
	ArrivalAirport     string         `db:"arrival_airport"`
	Status             string         `db:"status"`
	AircraftCode       string         `db:"aircraft_code"`
	ActualDeparture    sql.NullString `db:"actual_departure"`
	ActualArrival      sql.NullString `db:"actual_arrival"`
}

type FlightSeatInfo struct {
	FlightId           int     `db:"flight_id"`
	AircraftCode       string  `db:"aircraft_code"`
	EconomyAmount      int     `db:"economy_amount"`
	EconomyTotalPrice  float64 `db:"economy_total_price"`
	ComfortAmount      int     `db:"comfort_amount"`
	ComfortTotalPrice  float64 `db:"comfort_total_price"`
	BusinessAmount     int     `db:"business_amount"`
	BusinessTotalPrice float64 `db:"business_total_price"`
}
