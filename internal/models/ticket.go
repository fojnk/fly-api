package models

type Ticket struct {
	TicketNo      string `db:"ticket_no"`
	BookRef       string `db:"book_ref"`
	PassengerId   string `db:"passenger_id"`
	PassengerName string `db:"passenger_name"`
	ContactData   string `db:"contact_data"`
}
