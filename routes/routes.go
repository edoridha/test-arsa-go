package routes

import (
	"test-ara/controllers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Post(("/api/production-house"), controllers.CreateProductionHouse)
	r.Get(("/api/production-house"), controllers.GetProductionHouses)
	r.Get(("/api/production-house/{id}"), controllers.GetProductionHouseByID)
	r.Put(("/api/production-house/{id}"), controllers.UpdateProductionHouse)
	r.Delete(("/api/production-house/{id}"), controllers.DeleteProductionHouse)

	r.Get(("/api/film"), controllers.GetFilms)
	r.Post(("/api/film"), controllers.CreateFilm)
	r.Get(("/api/film/{id}"), controllers.GetFilmByID)
}