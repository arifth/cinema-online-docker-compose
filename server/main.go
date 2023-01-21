package main

import (
	"final-task/database"
	"final-task/pkg/mysql"
	"final-task/routes"
	"fmt"

	// "gorm-imp/routes"
	"net/http"

	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// stup for allowed Headers , and origin for CORS

var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})
var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	var port = os.Getenv("PORT")

	// initiate DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server jalan di alamat localhost:5000")
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
