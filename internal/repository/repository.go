package repository

import (
	"flyAPI/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	flightsTable     = "flights"
	airportDataTable = "airports_data"
	routeTable       = "routes"
)

type IAirRepository interface {
	GetAllSrcAndDestCities() (Cities, error)
	GetAllSrcAndDestAirports() (Airports, error)
	GetAirportsInCity(cityName string) ([]models.Airport, error)
	GetAirportByName(airportName string) (models.Airport, error)
}

type IScheduleRepository interface {
	GetInboundScheduleForAirport(airport string) ([]InboundSchedule, error)
	GetOutboundScheduleForAirport(airport string) ([]OutboundSchedule, error)
}

type IRouteRepository interface {
	GetRoutesFromAirport(aiport string) ([]models.Flight, error)
}

type Respository struct {
	IAirRepository
	IScheduleRepository
	IRouteRepository
}

func NewRepository(db *sqlx.DB) *Respository {
	return &Respository{
		IAirRepository:      NewAirRepo(db),
		IScheduleRepository: NewScheduleRepo(db),
		IRouteRepository:    NewRouteRepo(db),
	}
}
