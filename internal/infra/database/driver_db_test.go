package database_test

import (
	"fmt"
	"testing"

	"github.com/deduardo/gobrax/internal/entity"
	"github.com/deduardo/gobrax/internal/infra/database"
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

	err = db.AutoMigrate(&entity.Driver{})
	if err != nil {
		fmt.Println("failed to migrate database")
	}

	return db
}

func TestCreateDriver(t *testing.T) {
	db := setupDatabase(t)
	repo := database.NewDriver(db)

	driver := entity.Driver{
		Name:     "Edu",
		Email:    "edu@email.com",
		Password: "123456",
	}

	err := repo.Create(&driver)
	assert.NoError(t, err)
}

func TestFindDriverByEmail(t *testing.T) {
	db := setupDatabase(t)
	repo := database.NewDriver(db)

	expected := entity.Driver{
		Name:     "Edu",
		Email:    "john@example.com",
		Password: "123456",
	}

	err := db.Create(&expected).Error
	assert.NoError(t, err)

	driver, err := repo.FindByEmail(expected.Email)
	assert.NoError(t, err)
	assert.NotNil(t, driver)
	assert.Equal(t, expected.Email, driver.Email)
}
func TestFindDriverByID(t *testing.T) {
	db := setupDatabase(t)
	repo := database.NewDriver(db)

	expected := entity.Driver{
		Name:     "Edu",
		Email:    "john@example.com",
		Password: "123456",
	}

	err := db.Create(&expected).Error
	assert.NoError(t, err)

	driver, err := repo.FindByID(expected.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, driver)
	assert.Equal(t, expected.ID, driver.ID)
}

func TestGetAllDrivers(t *testing.T) {
	db := setupDatabase(t)
	repo := database.NewDriver(db)

	expected1 := entity.Driver{ID: uuid.New(), Name: "Driver 1", Email: "driver1@email.com", Password: "123456"}
	expected2 := entity.Driver{ID: uuid.New(), Name: "Driver 2", Email: "driver2@email.com", Password: "123"}
	expected3 := entity.Driver{ID: uuid.New(), Name: "Driver 3", Email: "driver3@email.com", Password: "123789"}

	_ = db.Create(&expected1)
	_ = db.Create(&expected2)
	_ = db.Create(&expected3)

	drivers, err := repo.GetAllDrivers()
	assert.NoError(t, err)
	assert.Len(t, drivers, 3)

	// Sua lógica de asserção continua aqui...
}

func TestUpdateDriver(t *testing.T) {
	db := setupDatabase(t)
	repo := database.NewDriver(db)

	driver := entity.Driver{ID: uuid.New(), Name: "Edu", Email: "edu@email.com", Password: "123456"}
	_ = db.Create(&driver)

	driver.Name = "Edu Updated"
	err := repo.Update(&driver)
	assert.NoError(t, err)

	var updatedDriver entity.Driver
	db.First(&updatedDriver, driver.ID)
	assert.Equal(t, "Edu Updated", updatedDriver.Name)
}

func TestDeleteDriver(t *testing.T) {
	db := setupDatabase(t)
	repo := database.NewDriver(db)

	driver := entity.Driver{ID: uuid.New(), Name: "Edu", Email: "edu@email.com", Password: "123456"}
	db.Create(&driver)

	err := repo.Delete(driver.ID.String())
	assert.NoError(t, err)

	var deletedDriver entity.Driver
	result := db.First(&deletedDriver, driver.ID)
	assert.Error(t, result.Error)
}
