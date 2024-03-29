package rest

import (
	"net/http"

	_ "github.com/plutus/banking-api/docs"       // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// @title        Plutus Banking API
// @version      1.0
// @description  This is a API representing a simple banking system.
// @host         localhost:3000
func NewRest(router *http.ServeMux, r Handler) error {

	// users
	router.HandleFunc("POST /user", r.CreateUser)
	router.HandleFunc("PUT /user/{id}", r.UpdateUser)
	router.HandleFunc("DELETE /user/{id}", r.DeleteUser)
	router.HandleFunc("GET /user/{id}", r.GetUser)
	router.HandleFunc("GET /users", r.GetUsers)

	// accounts
	router.HandleFunc("POST /user/{id}/account", r.CreateAccount)
	router.HandleFunc("GET /user/{user_id}/account/{account_id}", r.GetUserAccount)
	router.HandleFunc("DELETE /user/{user_id}/account/{account_id}", r.DeleteAccount)

	// transactions
	router.HandleFunc("POST /account/{id}/transaction", r.CreateTransaction)
	router.HandleFunc("GET /account/{id}/transactions", r.GetTransactions)

	router.HandleFunc("GET /healthz", r.Health)
	router.Handle("GET /swagger/*", httpSwagger.Handler())

	return nil
}
