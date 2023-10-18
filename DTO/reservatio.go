package dto

import (
	"time"
)

type ReservationDTO struct {
	ID       int64     `json:"id"`
	CheckIn  time.Time `json:"checkin" validate:"required" time_format:"2006-01-02"`
	CheckOut time.Time `json:"checkout" validate:"required" time_format:"2006-01-02"`
	IdHotel  int64     `json:"idHotel" validate:"required"`
	EMail    string    `json:"email" validate:"required,email"`
}

type ReservationsDTO []ReservationDTO
