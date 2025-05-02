package repository

import (
	"flyAPI/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BookingRepo struct {
	db *sqlx.DB
}

func NewBookingRepo(db *sqlx.DB) *BookingRepo {
	return &BookingRepo{
		db: db,
	}
}

func (p *BookingRepo) FindBookingByBookingRef(bookingRef string) (models.Booking, error) {
	var booking models.Booking

	query := fmt.Sprintf(`SELECT * from %s b WHERE b.book_ref = $1`, BookingTable)
	err := p.db.Get(&booking, query, bookingRef)
	return booking, err
}

func (p *BookingRepo) AddBooking(newBooking models.Booking) error {
	query := fmt.Sprintf(`INSERT INTO %s (book_ref, book_date, total_amount) VALUES ($1, $2, $3)`, BookingTable)

	_, err := p.db.Exec(query, newBooking.BookRef, newBooking.BookDate, newBooking.TotalAmount)
	return err
}
