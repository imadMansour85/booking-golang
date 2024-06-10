package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(body)
		return
	}
	log.Printf("Request Body: %s", body)
	// service.CreateBooking()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category")
}
