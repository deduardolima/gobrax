package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrIDIsRequired       = errors.New("id is required")
	ErrInvalidID          = errors.New("invalid id")
	ErrBrandIsRequired    = errors.New("brand is required")
	ErrModelIsRequired    = errors.New("model is required")
	ErrYearIsRequired     = errors.New("year is required and must be greater than 1886")
	ErrDriverIDIsRequired = errors.New("driverid is required")
	ErrInvalidDriverID    = errors.New("invalid driverid")
)

type Vehicle struct {
	ID       uuid.UUID `json:"id"`
	Brand    string    `json:"brand"`
	Model    string    `json:"model"`
	Year     int       `json:"year"`
	DriverID uuid.UUID `json:"driver_id"`
	CreateAt time.Time `json:"created_at"`
}

func NewVehicle(brand, model string, year int, driverID uuid.UUID) (*Vehicle, error) {
	if driverID == uuid.Nil {
		return nil, ErrDriverIDIsRequired
	}

	vehicle := &Vehicle{
		ID:       uuid.New(),
		Brand:    brand,
		Model:    model,
		Year:     year,
		DriverID: driverID,
	}
	err := vehicle.Validate()
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (v *Vehicle) Validate() error {
	if v.ID == uuid.Nil {
		return ErrIDIsRequired
	}
	if v.Brand == "" {
		return ErrBrandIsRequired
	}
	if v.Model == "" {
		return ErrModelIsRequired
	}
	if v.Year <= 1970 {
		return ErrYearIsRequired
	}
	return nil
}
