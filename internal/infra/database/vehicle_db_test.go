package database

import (
	"fmt"
	"testing"

	"github.com/deduardo/gobrax/internal/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}

	err = db.AutoMigrate(&entity.Vehicle{})
	if err != nil {
		fmt.Println("failed to migrate database")
	}

	return db
}

func TestCreateNewVehicle(t *testing.T) {
	db := setupDatabase(t)
	vehicle, err := entity.NewVehicle("Scania", "R620", 2012, uuid.New())
	assert.NoError(t, err)
	vehicleDB := NewVehicle(db)
	err = vehicleDB.Create(vehicle)
	assert.NoError(t, err)
	assert.NotEmpty(t, vehicle.ID)
}

func TestFindAllVehicles(t *testing.T) {
	db := setupDatabase(t)

	for i := 1; i < 24; i++ {
		vehicle, err := entity.NewVehicle(fmt.Sprintf("Scania %d", i), fmt.Sprintf("R6%d", i), 2012, uuid.New())
		assert.NoError(t, err)
		db.Create(vehicle)
	}
	vehicleDB := NewVehicle(db)
	vehicles, err := vehicleDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, vehicles, 10)
	assert.Equal(t, "Scania 1", vehicles[0].Brand)
	assert.Equal(t, "Scania 10", vehicles[9].Brand)

	vehicles, err = vehicleDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, vehicles, 10)
	assert.Equal(t, "Scania 11", vehicles[0].Brand)
	assert.Equal(t, "Scania 20", vehicles[9].Brand)

	vehicles, err = vehicleDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, vehicles, 3)
	assert.Equal(t, "Scania 21", vehicles[0].Brand)
	assert.Equal(t, "Scania 23", vehicles[2].Brand)
}

func TestFindVehicleByID(t *testing.T) {
	db := setupDatabase(t)
	vehicle, err := entity.NewVehicle("Scania", "R620", 2012, uuid.New())
	assert.NoError(t, err)
	db.Create(vehicle)
	vehicleDB := NewVehicle(db)
	foundVehicle, err := vehicleDB.FindByID(vehicle.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Scania", foundVehicle.Brand)
}

func TestUpdateVehicle(t *testing.T) {
	db := setupDatabase(t)
	vehicle, err := entity.NewVehicle("Scania", "R620", 2012, uuid.New())
	assert.NoError(t, err)
	db.Create(vehicle)
	vehicleDB := NewVehicle(db)
	vehicle.Brand = "Volvo"

	err = vehicleDB.Update(vehicle)
	assert.NoError(t, err)
	vehicle, err = vehicleDB.FindByID(vehicle.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Volvo", vehicle.Brand)
}

func TestDeleteVehicle(t *testing.T) {
	db := setupDatabase(t)
	vehicle, err := entity.NewVehicle("Scania", "R620", 2012, uuid.New())
	assert.NoError(t, err)
	db.Create(vehicle)
	vehicleDB := NewVehicle(db)
	err = vehicleDB.Delete(vehicle.ID.String())
	assert.NoError(t, err)
	_, err = vehicleDB.FindByID(vehicle.ID.String())
	assert.Error(t, err)

}
