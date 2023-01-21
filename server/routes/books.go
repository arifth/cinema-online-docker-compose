package routes

import (
	"final-task/handlers"
	"final-task/pkg/middleware"
	"final-task/pkg/mysql"
	"final-task/repositories"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	BookRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerBook(BookRepository)

	// NOTES: undefined routes will cause error method not allowed in client like postman

	r.HandleFunc("/books", h.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", h.GetBook).Methods("GET")
	r.HandleFunc("/book", middleware.Auth(middleware.UploadFile(h.CreateBook))).Methods("POST")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	// r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTrans))).Methods("PATCH")
	// r.HandleFunc("/trip/{id}", middleware.Auth(h.DeleteTrip)).Methods("DELETE")
}
