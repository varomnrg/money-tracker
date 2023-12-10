package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	handler "github.com/varomnrg/money-tracker/handler/category"
	mock_service "github.com/varomnrg/money-tracker/mocks/service/category"
	"github.com/varomnrg/money-tracker/model"
	service "github.com/varomnrg/money-tracker/service/category"
	"go.uber.org/mock/gomock"
)

func SetupCategoryHandler(t *testing.T) (*gomock.Controller, *handler.CategoryHandler, *mock_service.MockICategoryService) {
	ctrl := gomock.NewController(t)
	mockService := mock_service.NewMockICategoryService(ctrl)
	categoryHandler := handler.NewCategoryHandler(mockService)

	return ctrl, categoryHandler, mockService
}

func TestGetCategories_Handler(t *testing.T) {
	ctrl, categoryHandler, mockService := SetupCategoryHandler(t)
	defer ctrl.Finish()

	// Mock GetCategories
	mockService.EXPECT().GetCategories().Return(
		[]model.Category{
			{ID: "category-1", Name: "category1"},
			{ID: "category-2", Name: "category2"},
		},
		nil,
	)

	t.Run("Get Categories", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/categories", nil)

		recorder := httptest.NewRecorder()

		categoryHandler.GetCategories(recorder, req, nil)
		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")

		var categories []model.Category
		json.Unmarshal(recorder.Body.Bytes(), &categories)

		assert.Len(t, categories, 2, "Expected one category returned")
		assert.Equal(t, "category1", categories[0].Name, "Expected category1 in response array at index 0")
	})
}

func TestGetUserCategories_Handler(t *testing.T) {
	ctrl, categoryHandler, mockService := SetupCategoryHandler(t)
	defer ctrl.Finish()

	// Mock GetUserCategories
	mockService.EXPECT().GetUserCategories("user-1").Return(
		[]model.Category{
			{ID: "category-1", Name: "category1"},
			{ID: "category-2", Name: "category2"},
		},
		nil,
	)

	// Mock GetUserCategories with invalid id
	mockService.EXPECT().GetUserCategories("invalid_id").Return([]model.Category{}, service.ErrUserNotFound)
	

	t.Run("Get User Categories", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/categories/user-1", nil)

		recorder := httptest.NewRecorder()

		categoryHandler.GetUserCategories(recorder, req, nil)
		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")

		var categories []model.Category
		json.Unmarshal(recorder.Body.Bytes(), &categories)

		assert.Len(t, categories, 2, "Expected one category returned")
		assert.Equal(t, "category1", categories[0].Name, "Expected category1 in response array at index 0")
	})

	t.Run("Get User Categories with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/categories/invalid_id", nil)
		recorder := httptest.NewRecorder()

		categoryHandler.GetUserCategories(recorder, req, nil)

		assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected status Not Found")
	})
}

func TestGetCategory_Handler(t *testing.T) {
	ctrl, categoryHandler, mockService := SetupCategoryHandler(t)
	defer ctrl.Finish()

	// Mock GetCategory
	mockService.EXPECT().GetCategory("category-1").Return(model.Category{ID: "category-1", Name: "category1"}, nil)
	mockService.EXPECT().GetCategory("invalid_id").Return(model.Category{}, service.ErrCategoryNotFound)

	t.Run("Get Category", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/categories/category-1", nil)

		recorder := httptest.NewRecorder()

		categoryHandler.GetCategory(recorder, req, []httprouter.Param{{Key: "id", Value: "category-1"}})

		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")

		var category model.Category
		json.Unmarshal(recorder.Body.Bytes(), &category)

		assert.Equal(t, "category1", category.Name, "Expected category1 in response")
	})

	t.Run("Get Category with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/categories/invalid_id", nil)
		recorder := httptest.NewRecorder()

		categoryHandler.GetCategory(recorder, req, []httprouter.Param{{Key: "id", Value: "invalid_id"}})

		assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected status Not Found")
	})
}

func TestCreateCategory_Handler(t *testing.T) {
	ctrl, categoryHandler, mockService := SetupCategoryHandler(t)
	defer ctrl.Finish()

	// Mock CreateCategory
	mockService.EXPECT().CreateCategory("user-1", model.CategoryRequest{Name: "category1"}).Return(nil)

	t.Run("Create Category", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/categories", nil)

		recorder := httptest.NewRecorder()

		categoryHandler.CreateCategory(recorder, req, []httprouter.Param{{Key: "user_id", Value: "user-1"}})

		assert.Equal(t, http.StatusCreated, recorder.Code, "Expected status Created")
	})
}

func TestDeleteCategory_Handler(t *testing.T) {
	ctrl, categoryHandler, mockService := SetupCategoryHandler(t)
	defer ctrl.Finish()

	// Mock DeleteCategory
	mockService.EXPECT().DeleteCategory("category-1").Return(nil)

	// Mock DeleteCategory with invalid id
	mockService.EXPECT().DeleteCategory("invalid_id").Return(service.ErrCategoryNotFound)

	t.Run("Delete Category", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/categories/category-1", nil)

		recorder := httptest.NewRecorder()

		categoryHandler.DeleteCategory(recorder, req, []httprouter.Param{{Key: "id", Value: "category-1"}})

		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")
	})

	t.Run("Delete Category with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/categories/invalid_id", nil)
		recorder := httptest.NewRecorder()

		categoryHandler.DeleteCategory(recorder, req, []httprouter.Param{{Key: "id", Value: "invalid_id"}})

		assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected status Not Found")
	})
}