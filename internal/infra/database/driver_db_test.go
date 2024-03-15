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

func TestFindAllDrivers(t *testing.T) {
	db := setupDatabase(t)

	for i := 1; i < 24; i++ {
		driver, err := entity.NewDriver(fmt.Sprintf("Motorista%d", i), fmt.Sprintf("email%d@example.com", i), "123456")
		assert.NoError(t, err)
		result := db.Create(driver)
		assert.NoError(t, result.Error)
	}
	driverDB := database.NewDriver(db)
	drivers, err := driverDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, drivers, 10)
	assert.Equal(t, "Motorista1", drivers[0].Name)
	assert.Equal(t, "Motorista10", drivers[9].Name)

	drivers, err = driverDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, drivers, 10)
	assert.Equal(t, "Motorista11", drivers[0].Name)
	assert.Equal(t, "Motorista20", drivers[9].Name)

	drivers, err = driverDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, drivers, 3)
	assert.Equal(t, "Motorista21", drivers[0].Name)
	assert.Equal(t, "Motorista23", drivers[2].Name)
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
