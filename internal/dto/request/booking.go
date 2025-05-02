package request

type BookingRaceRequest struct {
	FlightsIds       []int
	FareCondition    string
	PassengerId      string
	PassengerName    string
	PassengerContact string
}

type BookingOneRaceRequest struct {
	FlightId         int
	FareCondition    string
	PassengerId      string
	PassengerName    string
	PassengerContact string
}
