package request

type FlightParams struct {
	Src            string
	Dest           string
	LenghtLimit    int
	FareConditions string
	DepartureDate  string
}
