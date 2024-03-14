package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDriver(t *testing.T) {
	driver, err := NewDriver("John Doe", "j@j.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, driver)
	assert.NotEmpty(t, driver.ID)
	assert.NotEmpty(t, driver.Password)
	assert.Equal(t, "John Doe", driver.Name)
	assert.Equal(t, "j@j.com", driver.Email)

}

func TestNewDriver_ValidatePassword(t *testing.T) {
	driver, err := NewDriver("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.True(t, driver.ValidatePassword("123456"))
	assert.False(t, driver.ValidatePassword("1234567"))
}
