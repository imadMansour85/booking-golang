package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/imadMansour85/bookings-app/internal/models"
	"github.com/imadMansour85/bookings-app/internal/service"
)

func JsonFormatter(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return fmt.Errorf("failed to parse JSON data: %v", err)
	}
	return nil
}
func BookingHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		var booking models.Booking
		if err := JsonFormatter(r, &booking); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		service.CreateBooking(booking)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(booking)

	case http.MethodGet:
		bookingsList := service.ListBookings()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bookingsList)
	}

	// service.CreateBooking()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category")
}
