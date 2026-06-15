package domain

import "time"

type Show struct {
	ID        string    `json:"id"`
	MovieID   string    `json:"movie_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
