package main

import (
	"net/http"

	"github.com/deduardo/gobrax/configs"
	"github.com/deduardo/gobrax/internal/entity"
	"github.com/deduardo/gobrax/internal/infra/database"
	"github.com/deduardo/gobrax/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Vehicle{}, &entity.Driver{})
	vehicleDB := database.NewVehicle(db)
	vehicleHandler := handlers.NewVehicleHandler(vehicleDB)
	driverDB := database.NewDriver(db)
	driverHandler := handlers.NewDriverHandler(driverDB, configs.TokenAuth, configs.JWTExperesIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/vehicles", vehicleHandler.CreateVehicle)
	r.Get("/vehicles", vehicleHandler.GetVehicles)
	r.Get("/vehicles/{id}", vehicleHandler.GetVehicleById)
	r.Put("/vehicles/{id}", vehicleHandler.UpdateVehicle)
	r.Delete("/vehicles/{id}", vehicleHandler.DeleteVehicle)

	//	r.Post("/drivers", driverHandler.CreateDriver)
	r.Post("/drivers/login", driverHandler.GetJWT)
	r.Get("/drivers/{id}", vehicleHandler.GetVehicleById)
	r.Put("/drivers/{id}", vehicleHandler.UpdateVehicle)
	r.Delete("/drivers/{id}", vehicleHandler.DeleteVehicle)

	http.ListenAndServe(":8080", r)

}
