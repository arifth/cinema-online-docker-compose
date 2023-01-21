package routes

import (
	"final-task/handlers"
	"final-task/pkg/middleware"
	"final-task/pkg/mysql"
	"final-task/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	FilmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(FilmRepository)

	r.HandleFunc("/films", h.FindFilm).Methods("GET")
	r.HandleFunc("/film/{id}", h.FindSingleFilm).Methods("GET")
	r.HandleFunc("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm))).Methods("POST")
	r.HandleFunc("/film/{id}", middleware.Auth(middleware.UploadFile(h.UpdateFilm))).Methods("PATCH")
	r.HandleFunc("/film/{id}", middleware.Auth(h.DeleteFilm)).Methods("DELETE")
}
