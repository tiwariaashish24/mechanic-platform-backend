package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"roadside-platform/internal/models"

	"github.com/gorilla/mux"
)

var bookings = []models.Booking{}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	booking.ID = strconv.Itoa(len(bookings) + 1)
	booking.Status = "pending"
	bookings = append(bookings, booking)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func AcceptBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, booking := range bookings {
		if booking.ID == id {
			bookings[i].Status = "accepted"
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bookings[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Booking not found"})
}
