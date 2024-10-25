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

func (uh *UserHandler) UserCreatePost(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(createdUser)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
