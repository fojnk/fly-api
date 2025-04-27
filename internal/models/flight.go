package models

type Flight struct {
	FlightId           int64  `db:"flight_id"`
	FlightNo           string `db:"flight_no"`
	ScheduledDeparture string `db:"scheduled_departure"`
	ScheduledArrival   string `db:"scheduled_arrival"`
	DepartureAirport   string `db:"departure_airport"`
	ArrivalAirport     string `db:"arrival_airport"`
	Status             string `db:"status"`
	AircraftCode       string `db:"aircraft_code"`
	ActualDeparture    string `db:"actual_departure"`
	ActualArrival      string `db:"actual_arrival"`
}
