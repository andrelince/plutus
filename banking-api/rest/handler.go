package rest

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/plutus/banking-api/pkg/slice"
	"github.com/plutus/banking-api/rest/definitions"
	"github.com/plutus/banking-api/rest/transformer"
	"github.com/plutus/banking-api/services"
)

type Handler struct {
	userConn    services.UserConnector
	accountConn services.AccountConnector
	validator   *validator.Validate
}

func NewHandler(
	userConn services.UserConnector,
	accountConn services.AccountConnector,
) Handler {
	return Handler{
		userConn:    userConn,
		accountConn: accountConn,
		validator: validator.New(
			validator.WithRequiredStructEnabled(),
		),
	}
}

// Healthz godoc
//
//  @Summary      Check service health
//  @Description  Check service health condition
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

	out, err := h.userConn.CreateUser(r.Context(), transformer.FromUserInputDefToEntity(user))
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
// @Router       /user/{id} [put]
//
// @Param        id    path  string                 true  "id of user to update"
// @Param        user  body  definitions.UserInput  true  "user to update"
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

	out, err := h.userConn.UpdateUser(r.Context(), uint(id), transformer.FromUserInputDefToEntity(user))
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
// @Router       /user/{id} [delete]
//
// @Param        id  path  string  true  "id of user to delete"
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
// @Router       /user/{id} [get]
//
// @Param        id  path  string  true  "id of user to retrieve"
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

// GetUsers godoc
//
// @Summary      Retrieve a list of users
// @Description  Retrieve a list of users in the system
// @Tags         user
// @Produce      json
// @Success      200  {array}  []definitions.User
// @Router       /users [get]
func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	out, err := h.userConn.GetUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	jsonOut, err := json.Marshal(
		slice.FromManyToMany(out, transformer.FromUserEntityToDef),
	)
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

// CreateAccount godoc
//
// @Summary      Create a user account
// @Description  Create a user account
// @Tags         account
// @Produce      json
// @Success      200  {object}  definitions.Account
// @Router       /user/{id}/account [post]
//
// @Param        id  path  string  true  "id of user for whom to create the account"
func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("user id is invalid"))
		return
	}

	out, err := h.accountConn.CreateAccount(r.Context(), uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	jsonOut, err := json.Marshal(transformer.FromAccountEntityToDef(out))
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

// GetUserAccount godoc
//
// @Summary      Retrieve a user account
// @Description  Retrieve a user account in the system
// @Tags         account
// @Produce      json
// @Success      200  {object}  definitions.Account
// @Router       /user/{user_id}/account/{account_id} [get]
//
// @Param        user_id     path  string  true  "id of user to retrieve"
// @Param        account_id  path  string  true  "id of account to retrieve"
func (h Handler) GetUserAccount(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil || userID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("user id is invalid"))
		return
	}

	accountID, err := strconv.Atoi(r.PathValue("account_id"))
	if err != nil || accountID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("account id is invalid"))
		return
	}

	out, err := h.accountConn.GetAccountByUserIDAndID(r.Context(), uint(userID), uint(accountID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	jsonOut, err := json.Marshal(transformer.FromAccountEntityToDef(out))
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

// CreateTransaction godoc
//
// @Summary      Create an account transaction
// @Description  Create an account transaction in the system
// @Tags         account
// @Produce      json
// @Success      200  {object}  definitions.Transaction
// @Router       /account/{id}/transaction [post]
//
// @Param        user  body  definitions.TransactionInput  true  "transaction to create"
// @Param        id    path  string                        true  "id of account to create transaction in"
func (h Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("account id is invalid"))
		return
	}

	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, errors.New("failed to read request body"))
		return
	}

	var in definitions.TransactionInput
	if err = json.Unmarshal(b, &in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteError(w, err)
		return
	}

	out, err := h.accountConn.CreateTransaction(r.Context(), uint(id), transformer.FromTransactionInputDefToModel(in))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		WriteError(w, err)
		return
	}

	jsonOut, err := json.Marshal(transformer.FromTransactionEntityToDef(out))
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
