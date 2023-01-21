package routes

import (
	"final-task/handlers"
	"final-task/pkg/mysql"
	"final-task/repositories"

	"github.com/gorilla/mux"
)

func userRoutes(r *mux.Router) {
	UserRepository := repositories.RepositoryUser(mysql.DB)

	h := handlers.HandlerUser(UserRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
	// r.HandleFunc("/userProfile/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProfile))).Methods("POST")
}
