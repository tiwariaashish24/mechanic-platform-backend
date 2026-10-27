package models

type Mechanic struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Specialty   string  `json:"specialty"`
	Location    string  `json:"location"`
	Rating      float64 `json:"rating"`
	IsAvailable bool    `json:"is_available"`
}

type Booking struct {
	ID           string `json:"id"`
	MechanicID   string `json:"mechanic_id"`
	CustomerName string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
	VehicleInfo  string `json:"vehicle_info"`
	Issue        string `json:"issue"`
	Status       string `json:"status"` // pending, accepted, completed, cancelled
	CreatedAt    string `json:"created_at"`
}
