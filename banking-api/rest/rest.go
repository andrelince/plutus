package rest

import (
	"net/http"

	"github.com/plutus/banking-api/pkg/logger"

	_ "github.com/plutus/banking-api/docs"       // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// @title        Plutus Banking API
// @version      1.0
// @description  This is a API representing a simple banking system.
// @host         locahost:3000
func NewRest(router *http.ServeMux, r Handler, log logger.Logger) error {

	// users
	router.HandleFunc("POST /user", r.CreateUser)

	router.HandleFunc("GET /healthz", r.Health)
	router.Handle("GET /swagger/*", httpSwagger.Handler())

	return nil
}
