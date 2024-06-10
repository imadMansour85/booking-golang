package repository

import "github.com/imadMansour85/bookings-app/internal/models"

var bookings []models.Booking

func SaveBooking(booking models.Booking) {
	bookings = append(bookings, booking)
}

func GetAllBookings() []models.Booking {
	return bookings
}
