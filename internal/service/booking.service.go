package service

import "booking-system/internal/repository"

type Booking struct {
	ID       string `json:"id"`
	Customer string `json:"customer"`
	Date     string `json:"date"`
}

func CreateBooking(booking Booking) {
	repository.SaveBooking(booking)
}
