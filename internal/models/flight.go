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
	FlightNo           string
	AircraftCode       string
	EconomyAmount      int
	EconomyTotalPrice  int
	ComfortAmount      int
	ComfortTotalPrice  int
	BusinessAmount     int
	BusinessTotalPrice int
}
