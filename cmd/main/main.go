package main

import (
	"booking-system/internal/handler"
	"booking-system/pkg/db"
	"log"
	"net/http"
)

func main() {
	db.Init() // Initialize the database connection
	http.HandleFunc("/booking", handler.BookingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
