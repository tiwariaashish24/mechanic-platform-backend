package controllers

import (
	"encoding/json"
	"net/http"

	"roadside-platform/internal/models"

	"github.com/gorilla/mux"
)

var mechanics = []models.Mechanic{
	{
		ID:          "1",
		Name:        "John Smith",
		Email:       "john@example.com",
		Phone:       "555-0101",
		Specialty:   "Engine Repair",
		Location:    "Downtown",
		Rating:      4.5,
		IsAvailable: true,
	},
	{
		ID:          "2",
		Name:        "Jane Doe",
		Email:       "jane@example.com",
		Phone:       "555-0102",
		Specialty:   "Brake Service",
		Location:    "Uptown",
		Rating:      4.8,
		IsAvailable: true,
	},
}

func GetAllMechanics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mechanics)
}

func GetMechanicByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, mechanic := range mechanics {
		if mechanic.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mechanic)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Mechanic not found"})
}
