package service

import (
	"flyAPI/internal/models"
	"flyAPI/internal/repository"
)

type AirService struct {
	airRepo repository.IAirRepository
}

func NewAirService(airRepo repository.IAirRepository) *AirService {
	return &AirService{
		airRepo: airRepo,
	}
}

func (a *AirService) GetAllCities() (repository.Cities, error) {
	return a.airRepo.GetAllSrcAndDestCities()
}

func (a *AirService) IsOriginExists(origin string) bool {
	return a.airRepo.IsOriginExists(origin)
}

func (a *AirService) GetAllAirports() (repository.Airports, error) {
	return a.airRepo.GetAllSrcAndDestAirports()
}

func (a *AirService) GetAirportsByCity(cityName string) ([]models.Airport, error) {
	return a.airRepo.GetAirportsInCity(cityName)
}
