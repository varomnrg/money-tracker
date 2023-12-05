package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/varomnrg/money-tracker/model"
	"github.com/varomnrg/money-tracker/service"
	"github.com/varomnrg/money-tracker/utils"
)

type UserHandler struct {
	service service.IUserService
}

var validate *validator.Validate

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{service: userService}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, "Unable to get all users", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	user, err := h.service.GetUser(id)
	if err != nil {

		if errors.Is(err, service.ErrUserNotFound) {
			utils.JSONError(w, err, http.StatusNotFound)
			return
		}

		http.Error(w, "Unable to get user", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := model.UserRequest{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unable to decode user", http.StatusInternalServerError)
		return
	}

	validate = validator.New()

	err = validate.Struct(user)
	if err != nil {
		errs := utils.MapErrors(err, validate)

		utils.JSONErrorMap(w, errs, http.StatusBadRequest)

		return
	}

	err = h.service.CreateUser(user)
	if err != nil {
		if errors.Is(err, service.ErrUsernameAlreadyExist) {
			utils.JSONError(w, err, http.StatusBadRequest)
			return
		}

		fmt.Println(err)

		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	user := model.UserRequest{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unable to decode user", http.StatusInternalServerError)
		return
	}

	validate = validator.New()

	err = validate.Struct(user)
	if err != nil {
		errs := utils.MapErrors(err, validate)

		utils.JSONErrorMap(w, errs, http.StatusBadRequest)

		return
	}

	err = h.service.UpdateUser(id, user)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			utils.JSONError(w, err, http.StatusNotFound)
			return
		}

		http.Error(w, "Unable to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated"))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	err := h.service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			utils.JSONError(w, err, http.StatusNotFound)
			return
		}

		http.Error(w, "Unable to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}
