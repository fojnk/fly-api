package models

type Seat struct {
	AircraftCode   string `db:"aircraft_code"`
	SeatNo         string `db:"seat_no"`
	FareConditions string `db:"fare_conditions"`
}
