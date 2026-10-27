package routers

import (
	"roadside-platform/internal/controllers"

	"github.com/gorilla/mux"
)

func SetRouter(router *mux.Router) {

	router.HandleFunc("/api/mechanics", controllers.GetAllMechanics).Methods("GET")
	router.HandleFunc("/api/mechanics/{id}", controllers.GetMechanicByID).Methods("GET")

	router.HandleFunc("/api/bookings", controllers.CreateBooking).Methods("POST")
	router.HandleFunc("/api/bookings/{id}/accept", controllers.AcceptBooking).Methods("POST")
}
