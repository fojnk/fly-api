package repository

import (
	"flyAPI/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TicketFlightsRepo struct {
	db *sqlx.DB
}

func NewTicketFlightsRepo(db *sqlx.DB) *TicketFlightsRepo {
	return &TicketFlightsRepo{
		db: db,
	}
}

func (t *TicketFlightsRepo) GetAllSoldSeatsByFlightAndAircraftCode(flightId, aircraftCode string) (models.FlightSeatInfo, error) {
	var info models.FlightSeatInfo

	query := fmt.Sprintf(`SELECT "tf.flight_id, tf.aircraft_code,
            SUM(CASE WHEN tf.fare_condition = 'Economy' THEN 1 ELSE 0 END) as economy_amount,
            SUM (CASE WHEN tf.fare_condition = 'Economy' THEN tf.amount ELSE null END) as economy_total_price,
            SUM(CASE WHEN tf.fare_condition = 'Business' THEN 1 ELSE 0 END) as business_amount,
            SUM (CASE WHEN tf.fare_condition = 'Business' THEN tf.amount ELSE null END) as business_total_price, 
            SUM(CASE WHEN tf.fare_condition = 'Comfort' THEN 1 ELSE 0 END) as comfort_amount, 
            SUM (CASE WHEN tf.fare_condition = 'Comfort' THEN tf.amount ELSE null END) as comport_total_price
            FROM %s tf WHERE tf.flight_id = $1 AND tf.aircraft_code = $2`, ticketFlightsTable)

	err := t.db.Get(&info, query, flightId, aircraftCode)
	return info, err
}

func (t *TicketFlightsRepo) AddTicketFlight(newTicketFlight models.TicketFlights) error {
	query := fmt.Sprintf(`INSERT INTO %s (ticket_no, flight_id, fare_conditions, amount) VALUES ($1, $2, $3, $4)`, ticketFlightsTable)

	_, err := t.db.Exec(query, newTicketFlight.TicketNo, newTicketFlight.FlightId, newTicketFlight.FareConditions, newTicketFlight.Amount)
	return err
}

func (t *TicketFlightsRepo) FindTicketFlight(ticketNo string) (models.TicketFlights, error) {
	var ticket models.TicketFlights

	query := fmt.Sprintf(`SELECT * from %s t WHERE t.ticket_no = $1 LIMIT 1`, ticketFlightsTable)
	err := t.db.Get(&ticket, query, ticketNo)
	return ticket, err
}
