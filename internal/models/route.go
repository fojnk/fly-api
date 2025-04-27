package models

import "time"

type Route struct {
	FlightNo             string        `db:"flight_no"`
	DepartureAirport     string        `db:"departure_airport"`
	DepartureAirportName string        `db:"departure_airport_name"`
	DepartureCity        string        `db:"departure_city"`
	ArrivalAirport       string        `db:"arrival_airport"`
	ArrivalAirportName   string        `db:"arrival_airport_name"`
	ArrivalCity          string        `db:"arrival_city"`
	Duration             time.Duration `db:"duration"`
	DaysOfWeek           []int         `db:"days_of_week"`
}
