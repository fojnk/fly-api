package models

type Airport struct {
	AirportCode string  `db:"airport_code" json:"airport_code"`
	AirportName string  `db:"airport_name" json:"airport_name"`
	City        string  `db:"city" json:"city"`
	Coordinates []uint8 `db:"coordinates" json:"coordinates"`
	Timezone    string  `db:"timezone" json:"timezone"`
}
