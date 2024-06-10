package service

import (
	"github.com/imadMansour85/bookings-app/internal/models"
	"github.com/imadMansour85/bookings-app/internal/repository"
)

func CreateBooking(booking models.Booking) {
	repository.SaveBooking(booking)
}

func ListBookings() []models.Booking {
	return repository.GetAllBookings()
}
