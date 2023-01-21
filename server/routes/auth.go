package routes

import (
	"final-task/handlers"
	"final-task/pkg/middleware"
	"final-task/pkg/mysql"
	"final-task/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	UserRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(UserRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	// NOTES: this routes function as checker everytime app js refreshed
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")

}
