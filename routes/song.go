package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func SongRoute(r *mux.Router){
	songRepository := repositories.RepositorySong(mysql.DB)
	h := handlers.HandlerSong(songRepository)


	r.HandleFunc("/createsong",middleware.Auth(middleware.UploadFile(middleware.UploadSong(h.CreateSong)))).Methods("POST")
	r.HandleFunc("/fsong",h.FindSong).Methods("GET")
	r.HandleFunc("/song/{id}",middleware.Auth(h.GetSongById)).Methods("GET")
}