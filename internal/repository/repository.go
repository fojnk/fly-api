package repository

import (
	"flyAPI/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	flightsTable        = "flights"
	airportDataTable    = "airports_data"
	routeTable          = "routes"
	seatsTable          = "seats"
	ticketTable         = "tickets"
	ticketFlightsTable  = "ticket_flights"
	BookingTable        = "bookings"
	BoardingPassesTable = "boarding_passes"
)

type IAirRepository interface {
	GetAllSrcAndDestCities() (Cities, error)
	GetAllSrcAndDestAirports() (Airports, error)
	GetAirportsInCity(cityName string) ([]models.Airport, error)
	GetAirportByNameOrCode(airportName string) (models.Airport, error)
	IsOriginExists(origin string) bool
}

type IScheduleRepository interface {
	GetInboundScheduleForAirport(airport string, time string) ([]InboundSchedule, error)
	GetOutboundScheduleForAirport(airport string, time string) ([]OutboundSchedule, error)
}

type IRouteRepository interface {
	GetRoutesFromAirport(aiport, startDate, endDate string) ([]models.Flight, error)
}

type ISeatRepository interface {
	GetSeatsByAircraftCode(aircraftCode string) (models.AircraftSeatsInfo, error)
	FindSeatsByAircraftCodeAndFareCondition(aircraftCode, fareConditions string) ([]models.Seat, error)
	FindSeatAmountByAircraftCodeAndFareCondition(aircraftCode, fareConditions string) (models.AircraftSeatsByFareCondition, error)
}

type ITicketRepository interface {
	FindTicketByTicketNo(ticketNo string) (models.Ticket, error)
	FindTicketsByBookRef(bookRef string) ([]models.Ticket, error)
	AddTicket(newTicket models.Ticket) error
}

type ITicketFlightsRepository interface {
	GetAllSoldSeatsByFlightAndAircraftCode(flightId int, aircraftCode string) (models.FlightSeatInfo, error)
	AddTicketFlight(newTicketFlight models.TicketFlights) error
	FindTicketFlight(ticketNo string) (models.TicketFlights, error)
}

type IFlightRepository interface {
	GetFlightByFlightId(flightId int) (models.Flight, error)
}

type IBookingRepo interface {
	FindBookingByBookingRef(bookingRef string) (models.Booking, error)
	AddBooking(newBooking models.Booking) error
}

type IBoardingPassRepo interface {
	FindLastBoardingNo(flightId int) (int, error)
	AddBoardingPass(newBoardingPass models.BoardingPass) error
	ExistsByFlightIdAndTicketNo(flightId int, ticketNo string) (int, error)
	FindBoardingPasses(flightId int) ([]models.BoardingPass, error)
}

type Respository struct {
	IAirRepository
	IScheduleRepository
	IRouteRepository
	ISeatRepository
	ITicketFlightsRepository
	ITicketRepository
	IFlightRepository
	IBookingRepo
	IBoardingPassRepo
}

func NewRepository(db *sqlx.DB) *Respository {
	return &Respository{
		IAirRepository:           NewAirRepo(db),
		IScheduleRepository:      NewScheduleRepo(db),
		IRouteRepository:         NewRouteRepo(db),
		ITicketRepository:        NewTicketRepo(db),
		ITicketFlightsRepository: NewTicketFlightsRepo(db),
		ISeatRepository:          NewSeatRepo(db),
		IFlightRepository:        NewFlightRepo(db),
		IBookingRepo:             NewBookingRepo(db),
		IBoardingPassRepo:        NewBoardingPassesRepo(db),
	}
}
