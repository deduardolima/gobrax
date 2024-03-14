package database

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/deduardo/gobrax/internal/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateDriver(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	assert.NoError(t, err)

	testDriverID := uuid.New()
	testDriver := entity.Driver{
		ID:       testDriverID,
		Name:     "Edu",
		Email:    "edu@email.com",
		Password: "123456",
	}

	// Mocka a chamada inicial para SELECT VERSION() feita pelo GORM
	mock.ExpectQuery(regexp.QuoteMeta("SELECT VERSION()")).WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("5.7"))

	// Configura as expectativas para início de transação, execução e commit
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `drivers` (`id`,`name`,`email`,`password`) VALUES (?,?,?,?)")).
		WithArgs(testDriver.ID, testDriver.Name, testDriver.Email, testDriver.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	driverRepo := NewDriver(gormDB)
	err = driverRepo.Create(&testDriver)
	assert.NoError(t, err)

	// Verifica se todas as expectativas foram atendidas
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

	// Aqui, você pode adicionar mais verificações se necessário

	var driverFound entity.Driver
	err = gormDB.First(&driverFound, "id = ?", testDriver.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, testDriver.ID, driverFound.ID)
	assert.Equal(t, testDriver.Name, driverFound.Name)
	assert.Equal(t, testDriver.Email, driverFound.Email)
}

func TestGetAllDrivers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to create gorm DB: %v", err)
	}

	mockDrivers := []entity.Driver{
		{ID: uuid.New(), Name: "Driver 1", Email: "driver1@example.com"},
		{ID: uuid.New(), Name: "Driver 2", Email: "driver2@example.com"},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(mockDrivers[0].ID, mockDrivers[0].Name, mockDrivers[0].Email).
		AddRow(mockDrivers[1].ID, mockDrivers[1].Name, mockDrivers[1].Email)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `drivers`")).WillReturnRows(rows)

	driverRepo := NewDriver(gormDB)
	drivers, err := driverRepo.GetAllDrivers()

	assert.NoError(t, err)
	assert.Len(t, drivers, 2)
	assert.Equal(t, mockDrivers, drivers)
}

func TestUpdateDriver(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to create gorm DB: %v", err)
	}

	mockDriver := entity.Driver{ID: uuid.New(), Name: "Driver Updated", Email: "driverupdated@example.com"}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `drivers` SET `name`=?,`email`=? WHERE `id`=?")).
		WithArgs(mockDriver.Name, mockDriver.Email, mockDriver.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	driverRepo := NewDriver(gormDB)
	err = driverRepo.Update(&mockDriver)

	assert.NoError(t, err)
}

func TestDeleteDriver(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	assert.NoError(t, err)

	mockDriverID := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `drivers` WHERE `drivers`.`id` = ?")).
		WithArgs(mockDriverID.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(mockDriverID.String(), "Edu", "edu@email.com", "123456"))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `drivers` WHERE `drivers`.`id` = ?")).
		WithArgs(mockDriverID.String()).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	driverRepo := NewDriver(gormDB)
	err = driverRepo.Delete(mockDriverID.String())

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPingDatabase(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT VERSION()")).WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7"))

	// Define a expectativa para a consulta de teste
	mock.ExpectQuery(regexp.QuoteMeta("SELECT 1")).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// Executa a consulta de teste
	var result int
	err = gormDB.Raw("SELECT 1").Scan(&result).Error
	assert.NoError(t, err)
	assert.Equal(t, 1, result)

	// Verifica se todas as expectativas foram cumpridas
	assert.NoError(t, mock.ExpectationsWereMet())
}
