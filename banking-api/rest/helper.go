package rest

import (
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}
