package entity

import (
	"time"

	"github.com/deduardo/gobrax/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type Driver struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	CreateAt time.Time `json:"created_at"`
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
