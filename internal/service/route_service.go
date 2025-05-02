package service

import (
	"flyAPI/internal/dto/request"
	"flyAPI/internal/models"
	"flyAPI/internal/repository"
	"slices"
	"time"

	"github.com/sirupsen/logrus"
)

type RouteService struct {
	routeRepo repository.IRouteRepository
	airRepo   repository.IAirRepository
}

func NewRouteService(routeRepo repository.IRouteRepository, airRepo repository.IAirRepository) *RouteService {
	return &RouteService{
		routeRepo: routeRepo,
		airRepo:   airRepo,
	}
}

func (s *RouteService) GetRoutes(cfg request.FlightParams) ([][]models.Flight, error) {
	routes := make([][]models.Flight, 0)

	visited := make([]string, 0)

	fromPoints, err := s.getPointsByOrigin(cfg.Src)
	if err != nil {
		return [][]models.Flight{}, err
	}

	logrus.Info("from points", fromPoints)

	toPoints, err := s.getPointsByOrigin(cfg.Dest)
	if err != nil {
		return [][]models.Flight{}, err
	}

	logrus.Info("to points", toPoints)

	nonDirectStartFlights := make([]models.Flight, 0)

	start, end, err := getStartEndDepartureDates(cfg.DepartureDate)
	if err != nil {
		return routes, err
	}

	//direct
	for _, fromPoint := range fromPoints {
		potentialRoutes, err := s.routeRepo.GetRoutesFromAirport(fromPoint.AirportCode, start, end)
		if err != nil {
			return routes, err
		}

		logrus.Info("potential routes", potentialRoutes)

		for _, potRoute := range potentialRoutes {
			if isDest(toPoints, potRoute) {
				routes = append(routes, []models.Flight{potRoute})
			} else {
				nonDirectStartFlights = append(nonDirectStartFlights, potRoute)
			}
		}
	}

	//connected recursive
	if cfg.LenghtLimit > 0 {
		for _, point := range nonDirectStartFlights {
			s.dfsSearch(routes, []models.Flight{point}, toPoints, cfg.LenghtLimit, &visited)
		}
	}

	return routes, nil
}

func (s *RouteService) dfsSearch(routes [][]models.Flight, path []models.Flight, toPoints []models.Airport, connections int, visitedCities *[]string) {
	if connections <= len(path)-1 {
		return
	}

	lastPoint := path[len(path)-1]
	start, end, err := getStartEndDepartureDates(lastPoint.ScheduledArrival)
	if err != nil {
		return
	}

	potantialPaths, err := s.routeRepo.GetRoutesFromAirport(lastPoint.ArrivalAirport, start, end)
	if err != nil {
		return
	}

	lastAirport, err := s.airRepo.GetAirportByNameOrCode(lastPoint.ArrivalAirport)
	if err != nil {
		return
	}

	(*visitedCities) = append(*visitedCities, lastAirport.City)

	for _, point := range potantialPaths {
		newPath := slices.Clone(path)
		newPath = append(newPath, point)

		if isDest(toPoints, point) {
			routes = append(routes, newPath)
		} else {
			if cityNotVisited(*visitedCities, lastAirport.City) {
				s.dfsSearch(routes, newPath, toPoints, connections-1, visitedCities)
			}
		}
	}

	(*visitedCities) = (*visitedCities)[:len((*visitedCities))-1]
}

func getStartEndDepartureDates(date string) (string, string, error) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return "", "", err
	}

	startOfDay := t.Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	return startOfDay.Format(time.RFC3339), endOfDay.Format(time.RFC3339), nil
}

func cityNotVisited(visited []string, city string) bool {
	for _, currCity := range visited {
		if currCity == city {
			return false
		}
	}

	return true
}

func isDest(toPoints []models.Airport, point models.Flight) bool {
	for _, toPoint := range toPoints {
		if toPoint.AirportCode == point.ArrivalAirport {
			return true
		}
	}

	return false
}

func (s *RouteService) getPointsByOrigin(origin string) ([]models.Airport, error) {
	var fromPoints []models.Airport

	airport, err := s.airRepo.GetAirportByNameOrCode(origin)
	if err != nil {
		fromPoints, err = s.airRepo.GetAirportsInCity(origin)
		if err != nil {
			return fromPoints, err
		}
	} else {
		fromPoints = append(fromPoints, airport)
	}

	return fromPoints, err
}
