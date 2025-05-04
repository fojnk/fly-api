package service

import (
	"flyAPI/internal/dto/request"
	"flyAPI/internal/dto/response"
	"flyAPI/internal/models"
	"flyAPI/internal/repository"
)

type IAirService interface {
	GetAllCities() (repository.Cities, error)
	GetAllAirports() (repository.Airports, error)
	GetAirportsByCity(cityName string) ([]models.Airport, error)
	IsOriginExists(origin string) bool
}

type IScheduleService interface {
	GetInboundSchedule(airport string) ([]repository.InboundSchedule, error)
	GetOutboundSchedule(airport string) ([]repository.OutboundSchedule, error)
}

type IRouteService interface {
	GetRoutes(cfg request.FlightParams) ([][]models.Flight, error)
}

type IBookingService interface {
	CreateBooking(data request.BookingRaceRequest) ([]response.BookingResponse, error)
	CheckIn(data request.CheckInRequest) error
}

type Service struct {
	IAirService
	IScheduleService
	IRouteService
	IBookingService
}

func NewService(repos *repository.Respository) *Service {
	return &Service{
		IAirService:      NewAirService(repos.IAirRepository),
		IScheduleService: NewScheduleService(repos.IScheduleRepository),
		IRouteService:    NewRouteService(repos.IRouteRepository, repos.IAirRepository),
		IBookingService: NewBookingService(
			repos.IAirRepository,
			repos.IFlightRepository,
			repos.ISeatRepository,
			repos.ITicketFlightsRepository,
			repos.IBookingRepo,
			repos.ITicketRepository,
			repos.IBoardingPassRepo),
	}
}
