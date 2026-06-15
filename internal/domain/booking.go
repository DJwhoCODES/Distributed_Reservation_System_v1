package domain

import "time"

type BookingStatus string

const (
	BookingConfirmed BookingStatus = "confirmed"
	BookingCancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID        string        `json:"id"`
	ShowID    string        `json:"show_id"`
	SeatID    string        `json:"seat_id"`
	UserID    string        `json:"user_id"`
	Status    BookingStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
