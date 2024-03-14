package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deduardo/gobrax/internal/dto"
	"github.com/deduardo/gobrax/internal/entity"
	"github.com/deduardo/gobrax/internal/infra/database"
	entityPkg "github.com/deduardo/gobrax/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type VehicleHandler struct {
	VehicleDB database.VehicleInterface
}

func NewVehicleHandler(db database.VehicleInterface) *VehicleHandler {
	return &VehicleHandler{
		VehicleDB: db,
	}
}

func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle dto.CreateVehicleInput
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	v, err := entity.NewVehicle(vehicle.Brand, vehicle.Model, vehicle.Year, vehicle.DriverID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.VehicleDB.Create(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (h *VehicleHandler) GetVehicleById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	vehicle, err := h.VehicleDB.FindByID(id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)
}

func (h *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var vehicle entity.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vehicle.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.VehicleDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.VehicleDB.Update(&vehicle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *VehicleHandler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	_, err := h.VehicleDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.VehicleDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *VehicleHandler) GetVehicles(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	vehicles, err := h.VehicleDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}
