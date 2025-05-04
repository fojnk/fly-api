package request

type BookingRaceRequest struct {
	FlightsIds       []int  `json:"flight_ids"`
	FareCondition    string `json:"fare_condition"`
	PassengerId      string `json:"passenger_id"`
	PassengerName    string `json:"passenger_name"`
	PassengerContact string `json:"passenger_contact"`
}

type BookingOneRaceRequest struct {
	FlightId         int
	FareCondition    string
	PassengerId      string
	PassengerName    string
	PassengerContact string
}

type CheckInRequest struct {
	TicketNo string `json:"ticket_no"`
	FlightId int64  `json:"flight_id"`
}
