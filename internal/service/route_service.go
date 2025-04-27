package service

import (
	"flyAPI/internal/models"
	"flyAPI/internal/repository"
	"slices"
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

func (s *RouteService) GetRoutes(cfg repository.RouteParams) ([][]models.Flight, error) {
	routes := make([][]models.Flight, 0)

	visited := make([]string, 0)

	fromPoints, isCity, err := s.getPointsByOrigin(cfg.DepartureAirport)
	if err != nil {
		return [][]models.Flight{}, err
	}
	if isCity {
		visited = append(visited, cfg.DepartureAirport)
	}

	toPoints, isCity, err := s.getPointsByOrigin(cfg.Origin)
	if err != nil {
		return [][]models.Flight{}, err
	}
	if isCity {
		visited = append(visited, cfg.Origin)
	}

	nonDirectStartFlights := make([]models.Flight, 0)

	//direct
	for _, fromPoint := range fromPoints {
		potentialRoutes, err := s.routeRepo.GetRoutesFromAirport(fromPoint.AirportName)
		if err != nil {
			continue
		}

		for _, potRoute := range potentialRoutes {
			if isDest(toPoints, potRoute) {
				routes = append(routes, []models.Flight{potRoute})
			} else {
				nonDirectStartFlights = append(nonDirectStartFlights, potRoute)
			}
		}
	}

	//connected recursive
	for _, point := range nonDirectStartFlights {
		s.dfsSearch(routes, []models.Flight{point}, toPoints, cfg.LimitLength, visited)
	}

	return routes, nil
}

func (s *RouteService) dfsSearch(routes [][]models.Flight, path []models.Flight, toPoints []models.Airport, connections int, visitedCities []string) {
	lastPoint := path[len(path)-1]
	potantialPaths, _ := s.routeRepo.GetRoutesFromAirport(lastPoint.ArrivalAirport)

	for _, point := range potantialPaths {
		newPath := slices.Clone(path)
		newPath = append(newPath, point)

		if isDest(toPoints, point) {
			routes = append(routes, newPath)
		} else {
			if len(path)+1 <= connections {
				airport, _ := s.airRepo.GetAirportByName(lastPoint.ArrivalAirport)
				if cityNotVisited(visitedCities, airport.City) {
					newVisited := slices.Clone(visitedCities)
					newVisited = append(newVisited, airport.City)
					s.dfsSearch(routes, newPath, toPoints, connections, newVisited)
				}
			} else {
				break
			}
		}
	}
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
		if toPoint.AirportName == point.ArrivalAirport {
			return true
		}
	}

	return false
}

func (s *RouteService) getPointsByOrigin(origin string) ([]models.Airport, bool, error) {
	var fromPoints []models.Airport
	isCity := false

	from, err := s.airRepo.GetAirportByName(origin)
	if len(fromPoints) == 0 || err != nil {
		fromPoints, err = s.airRepo.GetAirportsInCity(origin)
		if err != nil && len(fromPoints) != 0 {
			isCity = true
		}
	} else {
		fromPoints = append(fromPoints, from)
	}

	return fromPoints, isCity, err
}
