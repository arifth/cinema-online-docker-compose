package routes

import (
	"final-task/handlers"
	"final-task/pkg/middleware"
	"final-task/pkg/mysql"
	"final-task/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	CategoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(CategoryRepository)

	// NOTES: undefined routes will cause error method not allowed in client like postman

	r.HandleFunc("/categories", h.FindCategories).Methods("GET")
	r.HandleFunc("/category/{id}", h.FindCategory).Methods("GET")
	r.HandleFunc("/category", middleware.Auth(h.CreateCategory)).Methods("POST")
	r.HandleFunc("/category/{id}", middleware.Auth(h.UpdateCategory)).Methods("PATCH")
	r.HandleFunc("/category/{id}", middleware.Auth(h.DeleteCategory)).Methods("DELETE")
}
