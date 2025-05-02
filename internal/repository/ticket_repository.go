package repository

import (
	"flyAPI/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TicketRepo struct {
	db *sqlx.DB
}

func NewTicketRepo(db *sqlx.DB) *TicketRepo {
	return &TicketRepo{
		db: db,
	}
}

func (p *TicketRepo) FindTicketByTicketNo(ticketNo string) (models.Ticket, error) {
	var ticket models.Ticket

	query := fmt.Sprintf(`SELECT * from %s t WHERE t.ticket_no = $1`, ticketTable)
	err := p.db.Get(&ticket, query, ticketNo)
	return ticket, err
}

func (p *TicketRepo) FindTicketsByBookRef(bookRef string) ([]models.Ticket, error) {
	var ticket []models.Ticket

	query := fmt.Sprintf(`SELECT * from %s t WHERE t.book_ref = $1`, ticketTable)
	err := p.db.Select(&ticket, query, bookRef)
	return ticket, err
}

func (p *TicketRepo) AddTicket(newTicket models.Ticket) error {
	query := fmt.Sprintf(`INSERT INTO %s (ticket_no, book_ref, passenger_id, passenger_name, contact_data) VALUES ($1, $2, $3, $4, $5)`, ticketTable)

	_, err := p.db.Exec(query, newTicket.TicketNo, newTicket.BookRef, newTicket.PassengerId, newTicket.PassengerName, newTicket.ContactData)
	return err
}
