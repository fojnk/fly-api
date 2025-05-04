package repository

import (
	"flyAPI/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BoardingPassesRepo struct {
	db *sqlx.DB
}

func NewBoardingPassesRepo(db *sqlx.DB) *BoardingPassesRepo {
	return &BoardingPassesRepo{
		db: db,
	}
}

func (p *BoardingPassesRepo) FindLastBoardingNo(flightId int) (int, error) {
	var boardingNo int

	query := fmt.Sprintf(`SELECT boarding_no FROM %s b WHERE b.flight_id = $1 ORDER BY b.boarding_no DESC LIMIT 1`, BoardingPassesTable)
	err := p.db.Get(&boardingNo, query, flightId)
	return boardingNo, err
}

func (p *BoardingPassesRepo) AddBoardingPass(newBoardingPass models.BoardingPass) error {
	query := fmt.Sprintf(`INSERT INTO %s (ticket_no, seat_no, boarding_no, flight_id) VALUES ($1, $2, $3)`, BoardingPassesTable)

	_, err := p.db.Exec(query, newBoardingPass.TicketNo, newBoardingPass.SeatNo, newBoardingPass.BoardingNo, newBoardingPass.FlightId)
	return err
}

func (p *BoardingPassesRepo) ExistsByFlightIdAndTicketNo(flightId int, ticketNo string) (int, error) {
	var boardingNo int

	query := fmt.Sprintf(`SELECT boarding_no FROM %s b WHERE b.flight_id = $1 AND b.flight_id = $2 ORDER BY b.boarding_no DESC LIMIT 1`, BoardingPassesTable)
	err := p.db.Get(&boardingNo, query, flightId)
	return boardingNo, err
}
