package database

import (
	"github.com/deduardo/gobrax/internal/entity"
)

type DriverInterface interface {
	Create(driver *entity.Driver) error
	FindByEmail(emailId string) (entity.Driver, error)
}
