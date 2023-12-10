package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	service "github.com/varomnrg/money-tracker/service/category"
)

type CategoryHandler struct {
	service service.ICategoryService
}

func NewCategoryHandler(categoryService service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{service: categoryService}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("not implemented")
}

func (h *CategoryHandler) GetUserCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("not implemented")
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("not implemented")
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("not implemented")
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("not implemented")
}
