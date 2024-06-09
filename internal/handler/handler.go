package handler

import (
	"booking-system/internal/service"
	"encoding/json"
	"net/http"
)

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var booking service.Booking
		err := json.NewDecoder(r.Body).Decode(&booking)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Call the booking service to create a new booking
		service.CreateBooking(booking)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(booking)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
