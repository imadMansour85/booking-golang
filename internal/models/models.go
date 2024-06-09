package models

type Booking struct {
	ID       string `json:"id"`
	Customer string `json:"customer"`
	Date     string `json:"date"`
}
