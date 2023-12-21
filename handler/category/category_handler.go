package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/varomnrg/money-tracker/model"
	service "github.com/varomnrg/money-tracker/service/category"
	"github.com/varomnrg/money-tracker/utils"
)

type CategoryHandler struct {
	service service.ICategoryService
}

var validate *validator.Validate

func NewCategoryHandler(categoryService service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{service: categoryService}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categories, err := h.service.GetCategories()
	if err != nil {
		http.Error(w, "Internal server error: unable to get all categories", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(categories)

	if err != nil {
		http.Error(w, "Internal server error: could not marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *CategoryHandler) GetUserCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")

	categories, err := h.service.GetUserCategories(userID)
	if err != nil {
		if err == service.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal server error: unable to get user categories", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, "Internal server error: could not marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	category, err := h.service.GetCategory(id)
	if err != nil {
		if err == service.ErrCategoryNotFound {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal server error: unable to get category", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(category)
	if err != nil {
		http.Error(w, "Internal server error: could not marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := model.CategoryRequest{}
	userID := ps.ByName("userID")

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Internal server error: unable to decode category", http.StatusInternalServerError)
		return
	}

	validate = validator.New()
	err = validate.Struct(category)
	if err != nil {
		errs := utils.MapErrors(err, validate)
		
		utils.JSONErrorMap(w, errs, http.StatusBadRequest)

		http.Error(w, "Internal server error: unable to validate category", http.StatusInternalServerError)
		return
	}

	err = h.service.CreateCategory(userID, category)
	if err != nil {
		if err == service.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err == service.ErrCategoryAlreadyExist {
			http.Error(w, "Category already exist", http.StatusBadRequest)
			return
		}

		http.Error(w, "Internal server error: unable to create category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Category created"))
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	err := h.service.DeleteCategory(id)
	if err != nil {
		if err == service.ErrCategoryNotFound {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal server error: unable to delete category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category deleted"))
}
