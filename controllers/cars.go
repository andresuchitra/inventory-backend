package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/andresuchitra/inventory-backend/models"
	"github.com/gorilla/mux"
)

// GetCars - return all cars
// GET - /cars
func GetCars(w http.ResponseWriter, r *http.Request) {
	cars := []models.Car{}

	models.DB.Find(&cars)

	respondJSON(w, http.StatusOK, cars)
}

// CreateCar - create new car
// POST - /cars
func CreateCar(w http.ResponseWriter, r *http.Request) {
	inputCar := models.Car{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&inputCar); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Save(&inputCar).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, inputCar)
}

// GetCar - get specific car by :id
// GET - /cars/:id
func GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := validateParam(vars, w)
	if id == -1 {
		return
	}

	car := findCar(id, w, r)
	if car == nil {
		return
	}
	respondJSON(w, http.StatusOK, car)
}

// DeleteCar - delete a car
// DELETE - /cars/:id
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := validateParam(vars, w)
	if id == -1 {
		return
	}

	car := findCar(id, w, r)
	if car == nil {
		return
	}

	if err := models.DB.Delete(&car).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, car.ID)
}

// UpdateCar - update a car
// PUT - /cars/:id
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inputCar := models.InputCarUpdate{}

	id := validateParam(vars, w)
	if id == -1 {
		return
	}

	car := findCar(id, w, r)
	if car == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&inputCar); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Model(&car).Omit("ID").Updates(inputCar).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, car)
}

// find or report not found when ID is not valid
func findCar(id int, w http.ResponseWriter, r *http.Request) *models.Car {
	data := models.Car{}

	if err := models.DB.First(&data, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}
