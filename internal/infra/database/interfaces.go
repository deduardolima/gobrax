package database

import (
	"github.com/deduardo/gobrax/internal/entity"
)

type DriverInterface interface {
	Create(driver *entity.Driver) error
	FindByEmail(emailId string) (entity.Driver, error)
}

type VehicleInterface interface {
	Create(vehicle *entity.Vehicle) error
	FindAll(page, limit int, sort string) ([]entity.Vehicle, error)
	FindByID(id string) (*entity.Vehicle, error)
	Update(vehicle *entity.Vehicle) error
	Delete(id string) error
}
