package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/ftransaction",middleware.Auth(h.FindTransaction)).Methods("GET")
	r.HandleFunc("/getbylogin",middleware.Auth(h.GetByLogin)).Methods("GET")
	r.HandleFunc("/createtransaction",middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction/{id}",middleware.Auth(h.GetTransactionById)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")

}