package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewVehicle(t *testing.T) {
	driverID, _ := uuid.NewUUID()
	vehicle, err := NewVehicle("Scania", "R620", 2012, driverID)

	assert.Nil(t, err)
	assert.NotNil(t, vehicle)
	assert.NotEmpty(t, vehicle.ID)
	assert.Equal(t, "Scania", vehicle.Brand)
	assert.Equal(t, "R620", vehicle.Model)
	assert.Equal(t, 2012, vehicle.Year)
	assert.Equal(t, driverID, vehicle.DriverID)
}

func TestVehicleWhenBrandIsRequired(t *testing.T) {
	driverID, _ := uuid.NewUUID()
	vehicle, err := NewVehicle("", "R620", 2012, driverID)
	assert.Nil(t, vehicle)
	assert.Equal(t, ErrBrandIsRequired, err)
}

func TestVehicleWhenModelIsRequired(t *testing.T) {
	driverID, _ := uuid.NewUUID()
	vehicle, err := NewVehicle("Scania", "", 2012, driverID)
	assert.Nil(t, vehicle)
	assert.Equal(t, ErrModelIsRequired, err)
}

func TestVehicleWhenYearIsRequired(t *testing.T) {
	driverID, _ := uuid.NewUUID()
	vehicle, err := NewVehicle("Scania", "R620", -1, driverID)
	assert.Nil(t, vehicle)
	assert.Equal(t, ErrYearIsRequired, err)
}

func TestVehicleWhenDriverIDIsRequired(t *testing.T) {
	vehicle, err := NewVehicle("Scania", "R620", 2012, uuid.Nil)
	assert.Nil(t, vehicle)
	assert.Equal(t, ErrDriverIDIsRequired, err)
}
