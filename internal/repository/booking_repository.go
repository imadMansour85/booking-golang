package repository

import "booking-system/internal/models"

var bookings []models.Booking

func SaveBooking(booking models.Booking) {
	bookings = append(bookings, booking)
}
