package domain

import "time"

type ShowSeat struct {
	ID         string    `json:"id"`
	ShowID     string    `json:"show_id"`
	RowLabel   string    `json:"row_label"`
	SeatNumber int       `json:"seat_number"`
	SeatLabel  string    `json:"seat_label"`
	CreatedAt  time.Time `json:"created_at"`
}
