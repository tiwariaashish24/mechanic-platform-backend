package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"roadside-platform/internal/routers"
)

func main() {
	router := mux.NewRouter()

	// Health check endpoint
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	// API routes
	routers.SetRouter(router)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"Mechanic Platform API is running"}`))
}
