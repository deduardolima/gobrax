package entity

import (
	"time"

	"github.com/deduardo/gobrax/pkg/entity"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Driver struct {
	ID        entity.ID `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (driver *Driver) BeforeCreate(tx *gorm.DB) (err error) {
	if driver.ID == uuid.Nil {
		driver.ID = entity.NewID()
	}
	return
}

func NewDriver(name, email, password string) (*Driver, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Driver{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u Driver) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
