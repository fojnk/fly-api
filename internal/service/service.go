package service

import (
	"flyAPI/internal/dto/request"
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

type Service struct {
	IAirService
	IScheduleService
	IRouteService
}

func NewService(repos *repository.Respository) *Service {
	return &Service{
		IAirService:      NewAirService(repos.IAirRepository),
		IScheduleService: NewScheduleService(repos.IScheduleRepository),
		IRouteService:    NewRouteService(repos.IRouteRepository, repos.IAirRepository),
	}
}
