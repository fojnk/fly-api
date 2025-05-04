package models

type BoardingPass struct {
	TicketNo   string `db:"ticket_no`
	SeatNo     string `db:"seat_no"`
	FlightId   int64  `db:"flight_id"`
	BoardingNo int64  `db:"boarding_no"`
}
