package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/plutus/banking-api/rest/definitions"
	"github.com/plutus/banking-api/rest/transformer"
	"github.com/plutus/banking-api/services"
)

type Handler struct {
	userConn  services.UserConnector
	validator *validator.Validate
}

func NewHandler(
	userConn services.UserConnector,
) Handler {
	return Handler{
		userConn: userConn,
		validator: validator.New(
			validator.WithRequiredStructEnabled(),
		),
	}
}

// Healthz godoc
//
//	@Summary      Check service health
//	@Description  Check service health condition
//	@Tags         health
//	@Produce      plain
//	@Success      200  {string}  string  "OK"
//	@Router       /healthz [get]
func (h Handler) Health(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte(`OK`)); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

// CreateUser godoc
//
// @Summary      Create a user
// @Description  Create a user in the system
// @Tags         user
// @Produce      json
// @Success      200  {object}  definitions.User
// @Router       /user [post]
//
// @Param        user  body  definitions.UserInput  true  "user to create"
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("failed to read request body"))
		return
	}

	var user definitions.UserInput
	if err = json.Unmarshal(b, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, err)
		return
	}

	out, err := h.userConn.CreateUser(r.Context(), transformer.FromUserInputDefToModel(user))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	jsonOut, err := json.Marshal(transformer.FromUserEntityToDef(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonOut); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// UpdateUser godoc
//
// @Summary      Update a user
// @Description  Update a user in the system
// @Tags         user
// @Produce      json
// @Success      200  {object}  definitions.User
// @Router       /user [put]
//
// @Param        id    query  string                 true  "id of user to update"
// @Param        user  body   definitions.UserInput  true  "user to update"
func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("user id is invalid"))
		return
	}

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("failed to read request body"))
		return
	}

	var user definitions.UserInput
	if err = json.Unmarshal(b, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, err)
		return
	}

	out, err := h.userConn.UpdateUser(r.Context(), uint(id), transformer.FromUserInputDefToModel(user))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	jsonOut, err := json.Marshal(transformer.FromUserEntityToDef(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonOut); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteUser godoc
//
// @Summary      Delete a user
// @Description  Delete a user from the system
// @Tags         user
// @Produce      json
// @Success      200
// @Router       /user [put]
//
// @Param        id  query  string  true  "id of user to delete"
func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("user id is invalid"))
		return
	}

	if err = h.userConn.DeleteUser(r.Context(), uint(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUser godoc
//
// @Summary      Retrieve a user
// @Description  Retrieve a user in the system
// @Tags         user
// @Produce      json
// @Success      200  {object}  definitions.User
// @Router       /user [get]
//
// @Param        id  query  string  true  "id of user to retrieve"
func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("user id is invalid"))
		return
	}

	out, err := h.userConn.GetUserByID(r.Context(), uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	jsonOut, err := json.Marshal(transformer.FromUserEntityToDef(out))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonOut); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
