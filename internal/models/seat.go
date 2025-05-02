package models

type Seat struct {
	AircraftCode   string `db:"aircraft_code"`
	SeatNo         int    `db:"seat_no"`
	FareConditions string `db:"fare_conditions"`
}

type AircraftSeatsInfo struct {
	AircraftCode   string `db:"aircraft_code"`
	SeatsAmount    int    `db:"seat_amount"`
	EconomyAmount  int    `db:"econom_amount"`
	ComportAmount  int    `db:"comfort_amount"`
	BusinessAmount int    `db:"business_amount"`
}

type AircraftSeatsByFareCondition struct {
	AircraftCode string `db:"aircraft_code"`
	Amount       int    `db:"amount"`
}
