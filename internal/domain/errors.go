package domain

import "errors"

var (
	ErrSeatAlreadyHeld   = errors.New("seat already held")
	ErrSeatAlreadyBooked = errors.New("seat already booked")

	ErrHoldNotFound = errors.New("hold not found")
	ErrInvalidHold  = errors.New("invalid hold")
	ErrUnauthorized = errors.New("unauthorized")

	ErrBookingNotFound = errors.New("booking not found")
)
