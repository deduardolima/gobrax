package database

import (
	"github.com/deduardo/gobrax/internal/entity"
	"gorm.io/gorm"
)

type Driver struct {
	DB *gorm.DB
}

func NewDriver(db *gorm.DB) *Driver {
	return &Driver{DB: db}
}

func (d *Driver) Create(driver *entity.Driver) error {
	return d.DB.Create(driver).Error
}

func (d *Driver) FindByEmail(email string) (*entity.Driver, error) {
	var driver entity.Driver
	if err := d.DB.Where("email = ?", email).First(&driver).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (v *Driver) FindByID(id string) (*entity.Driver, error) {
	var driver entity.Driver
	err := v.DB.Where("id = ?", id).Error
	return &driver, err
}
func (d *Driver) FindAll(page, limit int, sort string) ([]entity.Driver, error) {
	var drivers []entity.Driver
	var err error
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	orderQuery := "create_at " + sort

	if page != 0 && limit != 0 {
		err = d.DB.Limit(limit).Offset((page - 1) * limit).Order(orderQuery).Find(&drivers).Error
	} else {
		err = d.DB.Order(orderQuery).Find(&drivers).Error
	}
	return drivers, err
}

/*
	func (d *Driver) GetAllDrivers() ([]entity.Driver, error) {
		var drivers []entity.Driver
		if err := d.DB.Find(&drivers).Error; err != nil {
			return nil, err
		}
		return drivers, nil
	}
*/
func (d *Driver) Update(driver *entity.Driver) error {
	return d.DB.Save(driver).Error
}

func (d *Driver) Delete(id string) error {
	return d.DB.Delete(&entity.Driver{}, "id = ?", id).Error
}
