package handlers

import (
	"encoding/json"
	"errors"
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
	var user domain.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdUser, err := uh.UserApi.CreateUser(user)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateId) {
			http.Error(w, "Server failed to generate a unique id \nCould not create user", http.StatusInternalServerError)
			return
		} else if errors.Is(err, domain.ErrBadRequest) {
			http.Error(w, "Validation Failed.  \nCould not create user", http.StatusBadRequest)
			return
		}
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
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "users not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, "Could not get users", http.StatusInternalServerError)
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
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "user not found"}`, http.StatusNotFound)
			return
		}
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
		http.Error(w, "Invalid request payload \nCould not create user", http.StatusBadRequest)
		return
	} else if errors.Is(err, domain.ErrBadRequest) {
		http.Error(w, "Validation Failed. \nCould not create user", http.StatusBadRequest)
		return
	}

	updatedUser, err := uh.UserApi.UpdateUser(id, user)

	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "user not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, "Could not update user", http.StatusBadRequest)
		return
	}

	err = WriteJSON(w, http.StatusOK, updatedUser, nil)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (uh *UserHandler) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := uh.UserApi.DeleteUser(id)

	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			http.Error(w, `{"error": "user not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, "Could not delete user", http.StatusInternalServerError)
		return
	}

	err = WriteJSON(w, http.StatusNoContent, "Successfully deleted user", nil) //task : returns 200 instead of 204
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}

}
