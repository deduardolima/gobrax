package dto

import "github.com/google/uuid"

type CreateVehicleInput struct {
	Brand    string    `json:"brand"`
	Model    string    `json:"model"`
	Year     int       `json:"year"`
	DriverID uuid.UUID `json:"driver_id"`
}
