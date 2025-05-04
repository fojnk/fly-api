package service

import (
	"errors"
	"flyAPI/internal/dto/request"
	"flyAPI/internal/dto/response"
	"flyAPI/internal/models"
	"flyAPI/internal/repository"
	"math/rand"
	"strings"
	"time"
)

type BookingService struct {
	airRepo           repository.IAirRepository
	flightRepo        repository.IFlightRepository
	seatRepo          repository.ISeatRepository
	ticketFlightsRepo repository.ITicketFlightsRepository
	bookingRepo       repository.IBookingRepo
	ticketRepo        repository.ITicketRepository
	boardingPassRepo  repository.IBoardingPassRepo
}

func NewBookingService(
	airRepo repository.IAirRepository,
	flightRepo repository.IFlightRepository,
	seatRepo repository.ISeatRepository,
	ticketFlightsRepo repository.ITicketFlightsRepository,
	bookingRepo repository.IBookingRepo,
	ticketRepo repository.ITicketRepository,
	boardingPassRepo repository.IBoardingPassRepo) *BookingService {

	return &BookingService{
		airRepo:           airRepo,
		flightRepo:        flightRepo,
		seatRepo:          seatRepo,
		ticketFlightsRepo: ticketFlightsRepo,
		bookingRepo:       bookingRepo,
		ticketRepo:        ticketRepo,
		boardingPassRepo:  boardingPassRepo,
	}
}

func (b *BookingService) CheckIn(data request.CheckInRequest) error {
	_, err := b.ticketRepo.FindTicketByTicketNo(data.TicketNo)
	if err != nil {
		return err
	}

	if _, err = b.flightRepo.GetFlightByFlightId(int(data.FlightId)); err != nil {
		return err
	}

	if _, err = b.boardingPassRepo.ExistsByFlightIdAndTicketNo(int(data.FlightId), data.TicketNo); err == nil {
		return errors.New(`passenger already checked in`)
	}

	lastNo, err := b.boardingPassRepo.FindLastBoardingNo(int(data.FlightId))
	if err != nil {
		lastNo = 1
	}
	lastNo += 1

	newBoardingPass := models.BoardingPass{
		BoardingNo: int64(lastNo),
		FlightId:   data.FlightId,
		TicketNo:   data.TicketNo,
		SeatNo:     "10b",
	}

	err = b.boardingPassRepo.AddBoardingPass(newBoardingPass)
	return err
}

func (b *BookingService) CreateBooking(data request.BookingRaceRequest) ([]response.BookingResponse, error) {
	responses := make([]response.BookingResponse, len(data.FlightsIds))
	for _, flightId := range data.FlightsIds {
		response, err := b.BookOneRace(request.BookingOneRaceRequest{
			FlightId:         flightId,
			PassengerId:      data.PassengerId,
			PassengerName:    data.PassengerName,
			PassengerContact: data.PassengerContact,
			FareCondition:    data.FareCondition,
		})
		responses = append(responses, response)
		if err != nil {
			return responses, err
		}
	}

	return responses, nil
}

func (b *BookingService) BookOneRace(data request.BookingOneRaceRequest) (response.BookingResponse, error) {
	currFlight, err := b.flightRepo.GetFlightByFlightId(data.FlightId)
	if err != nil {
		return response.BookingResponse{}, err
	}

	if currFlight.ActualArrival.Valid || currFlight.ActualDeparture.Valid {
		return response.BookingResponse{}, errors.New("flight already finished")
	}

	seat, err := b.seatRepo.FindSeatsByAircraftCodeAndFareCondition(currFlight.AircraftCode, data.FareCondition)
	if err != nil || seat.Amount > 0 {
		return response.BookingResponse{}, errors.New("seats not found")
	}

	info, err := b.ticketFlightsRepo.GetAllSoldSeatsByFlightAndAircraftCode(currFlight.FlightNo, currFlight.AircraftCode)
	if err != nil {
		return response.BookingResponse{}, err
	}

	var totalPrice int64
	switch data.FareCondition {
	case "Economy":
		totalPrice = int64(info.EconomyTotalPrice) / int64(info.EconomyAmount)
	case "Comfort":
		totalPrice = int64(info.ComfortTotalPrice) / int64(info.ComfortAmount)
	case "Business":
		totalPrice = int64(info.BusinessTotalPrice) / int64(info.BusinessAmount)
	default:
		return response.BookingResponse{}, errors.New("unknown fare condition")
	}

	booking := models.Booking{
		BookDate:    time.Now().Format(time.RFC3339),
		TotalAmount: totalPrice,
		BookRef:     b.generateUniqueBookRef(),
	}

	if err := b.bookingRepo.AddBooking(booking); err != nil {
		return response.BookingResponse{}, err
	}

	ticket := models.Ticket{
		TicketNo:      b.generateUniqueTicketNo(),
		BookRef:       booking.BookRef,
		PassengerId:   data.PassengerId,
		PassengerName: data.PassengerName,
		ContactData:   data.PassengerContact,
	}

	if err := b.ticketRepo.AddTicket(ticket); err != nil {
		return response.BookingResponse{}, err
	}

	ticketFlights := models.TicketFlights{
		TicketNo:       ticket.TicketNo,
		FlightId:       currFlight.FlightId,
		FareConditions: data.FareCondition,
		Amount:         totalPrice,
	}

	if err := b.ticketFlightsRepo.AddTicketFlight(ticketFlights); err != nil {
		return response.BookingResponse{}, err
	}

	return response.BookingResponse{
		TicketNo: ticket.TicketNo,
		BookRef:  ticket.BookRef,
	}, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return strings.ToUpper(string(b))
}

func (b *BookingService) generateUniqueBookRef() string {
	bookingRefId := ""
	for {
		bookingRefId = randomString(6)
		if _, err := b.bookingRepo.FindBookingByBookingRef(bookingRefId); err != nil {
			break
		}
	}
	return bookingRefId
}

func (b *BookingService) generateUniqueTicketNo() string {
	ticketNo := ""
	for {
		ticketNo = randomString(13)
		if _, err := b.ticketRepo.FindTicketByTicketNo(ticketNo); err != nil {
			break
		}
	}
	return ticketNo
}
