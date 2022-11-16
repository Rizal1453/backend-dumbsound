package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router){
	UserRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(UserRepository)


	r.HandleFunc("/users",h.FindUsers).Methods("GET")
	r.HandleFunc("/createuser",h.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}",h.GetUserById).Methods("GET")
}