package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	ArtisRoute(r)
	SongRoute(r)
	TransactionRoutes(r)
}