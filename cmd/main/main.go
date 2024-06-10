package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imadMansour85/bookings-app/internal/handler"
	"github.com/imadMansour85/bookings-app/pkg/db"
)

func main() {
	db.Init() // Initialize the database connection
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/booking", handler.BookingHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
