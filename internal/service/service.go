package service

import (
	"flyAPI/internal/models"
	"flyAPI/internal/repository"
)

type IAirService interface {
	GetAllCities() (repository.Cities, error)
	GetAllAirports() (repository.Airports, error)
	GetAirportsByCity(cityName string) ([]models.Airport, error)
}

type IScheduleService interface {
	GetInboundSchedule(airport string) ([]repository.InboundSchedule, error)
	GetOutboundSchedule(airport string) ([]repository.OutboundSchedule, error)
}

type IRouteRepository interface {
}

type Service struct {
	IAirService
	IScheduleService
	IRouteRepository
}

func NewService(repos *repository.Respository) *Service {
	return &Service{
		IAirService:      NewAirService(repos.IAirRepository),
		IScheduleService: NewScheduleService(repos.IScheduleRepository),
		IRouteRepository: NewRouteService(repos.IRouteRepository, repos.IAirRepository),
	}
}
