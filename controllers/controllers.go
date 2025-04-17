package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"test-ara/db"
	"test-ara/models"

	"github.com/go-chi/chi/v5"
)

func CreateProductionHouse(w http.ResponseWriter, r *http.Request) {
	var productionHouse models.ProductionHouse
	if err := json.NewDecoder(r.Body).Decode(&productionHouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database := db.GetDBInstance()
	if err := database.Create(&productionHouse).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(productionHouse)
}

func GetProductionHouses(w http.ResponseWriter, r *http.Request) {
	database := db.GetDBInstance()
	var productionHouses []models.ProductionHouse
	if err := database.Find(&productionHouses).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productionHouses)
}

func GetProductionHouseByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("ID received:", id)
	database := db.GetDBInstance()
	var productionHouse models.ProductionHouse
	if err := database.First(&productionHouse, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productionHouse)
}

func UpdateProductionHouse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("ID received:", id)
	database := db.GetDBInstance()
	var productionHouse models.ProductionHouse
	if err := database.First(&productionHouse, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&productionHouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.Save(&productionHouse).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productionHouse)
}

func DeleteProductionHouse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("ID received:", id)
	database := db.GetDBInstance()
	var productionHouse models.ProductionHouse
	if err := database.First(&productionHouse, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := database.Delete(&productionHouse).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var film models.Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database := db.GetDBInstance()
	if err := database.Create(&film).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(film)
}

func GetFilms(w http.ResponseWriter, r *http.Request) {
	database := db.GetDBInstance()
	var films []models.Film
	if err := database.Find(&films).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(films)
}

func GetFilmByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("ID received:", id)
	database := db.GetDBInstance()
	var film models.Film
	if err := database.First(&film, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(film)
}

