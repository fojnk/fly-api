package models

type TicketFlights struct {
	TicketNo       string `json:"ticket_no" db:"ticket_no"`
	FlightId       int64  `json:"flight_id" db:"flight_id"`
	FareConditions string `json:"fare_conditions" db:"fare_conditions"`
	Amount         int64  `json:"amount" db:"amount"`
}
