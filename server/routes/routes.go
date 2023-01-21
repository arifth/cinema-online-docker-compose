package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	userRoutes(r)
	AuthRoutes(r)
	FilmRoutes(r)
	CategoryRoutes(r)
	BookRoutes(r)
}
