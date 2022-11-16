package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func ArtisRoute(r *mux.Router){
	artisRepository := repositories.RepositoryArtis(mysql.DB)
	h := handlers.HandlerArtis(artisRepository)

	r.HandleFunc("/fartis",middleware.Auth(h.FindArtis)).Methods("GET")
	r.HandleFunc("/artis/{id}",middleware.Auth(h.GetArtisById)).Methods("GET")
	r.HandleFunc("/createartis",middleware.Auth(h.CreateArtis)).Methods("POST")
}