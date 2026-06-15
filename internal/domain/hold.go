package domain

import "time"

type Hold struct {
	ID        string    `json:"id"`
	ShowID    string    `json:"show_id"`
	SeatID    string    `json:"seat_id"`
	UserID    string    `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
