package database

import (
	"github.com/deduardo/gobrax/internal/entity"
	"gorm.io/gorm"
)

type Vehicle struct {
	DB *gorm.DB
}

func NewVehicle(db *gorm.DB) *Vehicle {
	return &Vehicle{DB: db}
}

func (v *Vehicle) Create(vehicle *entity.Vehicle) error {
	return v.DB.Create(vehicle).Error
}

func (v *Vehicle) FindAll(page, limit int, sort string) ([]entity.Vehicle, error) {
	var vehicles []entity.Vehicle
	var err error
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	orderQuery := "created_at " + sort

	if page != 0 && limit != 0 {
		err = v.DB.Limit(limit).Offset((page - 1) * limit).Order(orderQuery).Find(&vehicles).Error
	} else {
		err = v.DB.Preload("Driver").Order(orderQuery).Find(&vehicles).Error
	}
	return vehicles, err
}

func (db *Vehicle) FindByID(id string) (*entity.Vehicle, error) {
	var vehicle entity.Vehicle
	if err := db.DB.Preload("Driver").First(&vehicle, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (v *Vehicle) Update(vehicle *entity.Vehicle) error {
	_, err := v.FindByID(vehicle.ID.String())
	if err != nil {
		return err
	}
	return v.DB.Save(vehicle).Error
}

func (v *Vehicle) Delete(id string) error {
	vehicle, err := v.FindByID(id)
	if err != nil {
		return err
	}
	return v.DB.Delete(vehicle).Error
}
