package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mahider-T/GoCrudChallange/internal/application/core/domain"
	"github.com/Mahider-T/GoCrudChallange/internal/ports/apiport"
)

type UserHandler struct {
	UserApi apiport.UserApiPort
}

func NewUserHandler(userApi apiport.UserApiPort) *UserHandler {
	return &UserHandler{
		UserApi: userApi,
	}
}

func (uh *UserHandler) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdUser, err := uh.UserApi.CreateUser(user)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	err = WriteJSON(w, http.StatusCreated, createdUser, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}

}

func (uh *UserHandler) UserGetAllHandler(w http.ResponseWriter, r *http.Request) {

	users, err := uh.UserApi.GetAllUsers()
	if err != nil {
		http.Error(w, "Could not get users", http.StatusInternalServerError) //task : handle empty list
	}
	err = WriteJSON(w, http.StatusOK, users, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (uh *UserHandler) UserGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := uh.UserApi.GetUserByID(id)
	if err != nil {
		http.Error(w, "Could not get user by id", http.StatusInternalServerError)
		return
	}

	err = WriteJSON(w, http.StatusOK, user, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (uh *UserHandler) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var user domain.UpdateUserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedUser, err := uh.UserApi.UpdateUser(id, user)

	if err != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	err = WriteJSON(w, http.StatusOK, updatedUser, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}

}
