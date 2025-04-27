package models

type Booking struct {
	BookRef     string `db:"book_ref"`
	BookDate    string `db:"book_date"`
	TotalAmount int64  `db:"total_amount"`
}
