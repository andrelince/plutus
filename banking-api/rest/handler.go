package rest

import "net/http"

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

// Healthz godoc
//
//  @Summary      Check service health
//  @Description  Check servivce health condition
//  @Tags         health
//  @Produce      plain
//  @Success      200  {string}  string  "OK"
//  @Router       /healthz [get]
func (h Handler) Health(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`OK`)); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
