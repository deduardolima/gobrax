package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/deduardo/gobrax/internal/dto"
	"github.com/deduardo/gobrax/internal/entity"
	"github.com/deduardo/gobrax/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type DriverHandler struct {
	DriverDB      database.DriverInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewDriverHandler(db database.DriverInterface, jwt *jwtauth.JWTAuth, JwtExperiesIn int) *DriverHandler {
	return &DriverHandler{
		DriverDB:      db,
		Jwt:           jwt,
		JwtExperiesIn: JwtExperiesIn,
	}
}

func (h *DriverHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var driver dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	d, err := h.DriverDB.FindByEmail(driver.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !d.ValidatePassword(driver.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": d.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExperiesIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *DriverHandler) CreateDriver(w http.ResponseWriter, r *http.Request) {
	var driver dto.CreateDriverInput
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	d, err := entity.NewDriver(driver.Name, driver.Email, driver.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.DriverDB.Create(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
