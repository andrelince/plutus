package rest

import (
	"net/http"

	"github.com/plutus/banking-api/pkg/logger"
)

func NewRest(router *http.ServeMux, r Handler, log logger.Logger) error {

	router.HandleFunc("GET /healthz", r.Health)

	return nil
}
