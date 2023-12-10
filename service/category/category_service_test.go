package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mockCatRepo "github.com/varomnrg/money-tracker/mocks/repository/category"
	mockUserRepo "github.com/varomnrg/money-tracker/mocks/repository/user"
	"github.com/varomnrg/money-tracker/model"
	service "github.com/varomnrg/money-tracker/service/category"
	"go.uber.org/mock/gomock"
)


func SetupCategoryService(t *testing.T) (*gomock.Controller, *service.CategoryService, *mockCatRepo.MockICategoryRepository,*mockUserRepo.MockIUserRepository){
	ctrl := gomock.NewController(t)
	mockUserRepo := mockUserRepo.NewMockIUserRepository(ctrl)
	mockCatRepo := mockCatRepo.NewMockICategoryRepository(ctrl)
	categoryService := service.NewCategoryService(mockCatRepo, mockUserRepo)

	return ctrl, categoryService, mockCatRepo, mockUserRepo
}

func TestGetCategories(t *testing.T) {
	ctrl, categoryService, mockCatRepo, _ := SetupCategoryService(t)
	defer ctrl.Finish()

	// Mock for GetCategories
	mockCatRepo.EXPECT().GetCategories().Return([]model.Category{
		{ID: "cat-1", Name: "category1", User_ID: "user-1"},
		{ID: "cat-2", Name: "category2", User_ID: "user-2"},
		}, 
		nil,
	)

	t.Run("Get Categories", func(t *testing.T){
		categories, err := categoryService.GetCategories()

		assert.NoError(t,err)
		assert.Len(t, categories, 2)		
		assert.Equal(t, "category1", categories[0].Name)
	})
}

func TestGetUserCategories(t *testing.T) {
	ctrl, categoryService, mockCatRepo, mockUserRepo := SetupCategoryService(t)
	defer ctrl.Finish()

	// Mock for GetUserCategories
	mockUserRepo.EXPECT().GetUser("user-1").Return(model.UserResponse{ID: "user-1", Username: "user1", Email: "user1@gmail.com", Created_At: time.Now()}, nil)
	mockCatRepo.EXPECT().GetUserCategories("user-1").Return([]model.Category{
		{ID: "cat-1", Name: "category1"},
		{ID: "cat-2", Name: "category2"},
		}, 
		nil,
	)

	// Mock for GetUserCategories without categories
	mockUserRepo.EXPECT().GetUser("user-2").Return(model.UserResponse{ID: "user-2", Username: "user2", Email: "user2@gmail.com", Created_At: time.Now()}, nil)
	mockCatRepo.EXPECT().GetUserCategories("user-2").Return([]model.Category{}, nil)

	// Mock for GetUserCategories with invalid user id
	mockUserRepo.EXPECT().GetUser("invalid_id").Return(model.UserResponse{}, nil)

	t.Run("Get User Categories", func(t *testing.T){
		categories, err := categoryService.GetUserCategories("user-1")

		assert.NoError(t,err)
		assert.Len(t, categories, 2)		
		assert.Equal(t, "category1", categories[0].Name)
	})

	t.Run("Get User Categories without categories", func(t *testing.T){
		categories, err := categoryService.GetUserCategories("user-2")

		assert.NoError(t,err)
		assert.Len(t, categories, 0)		
	})

	t.Run("Get User Categories with invalid user id", func(t *testing.T){
		_, err := categoryService.GetUserCategories("invalid_id")

		assert.ErrorIs(t, err, service.ErrUserNotFound)
	})
}

func TestGetCategory(t *testing.T) {
	ctrl, categoryService, mockCatRepo, _ := SetupCategoryService(t)
	defer ctrl.Finish()

	// Mock for GetCategory
	mockCatRepo.EXPECT().GetCategory("cat-1").Return(model.Category{ID: "cat-1", Name: "category1", User_ID: "user-1"}, nil)
	
	// Mock for GetCategory with invalid id
	mockCatRepo.EXPECT().GetCategory("invalid_id").Return(model.Category{}, nil)

	t.Run("Get Category", func(t *testing.T){
		category, err := categoryService.GetCategory("cat-1")

		assert.NoError(t,err)
		assert.Equal(t, "category1", category.Name)
	})

	t.Run("Get Category with invalid id", func(t *testing.T){
		_, err := categoryService.GetCategory("invalid_id")

		assert.ErrorIs(t, err, service.ErrCategoryNotFound)
	})
}

func TestCreateCategory(t *testing.T) {
	ctrl, categoryService, mockCatRepo, mockUserRepo := SetupCategoryService(t)
	defer ctrl.Finish()

	// Mock for Create Category
	mockUserRepo.EXPECT().GetUser("user-1").Return(model.UserResponse{ID: "user-1", Username: "user1", Email: "user1@gmail.com", Created_At: time.Now()}, nil)
	mockCatRepo.EXPECT().IsUserCategoryExist("user-1", "category1").Return(false)
	mockCatRepo.EXPECT().CreateCategory("user-1", gomock.Any()).Return(nil)

	// Mock for Create Category with existing category name
	mockUserRepo.EXPECT().GetUser("user-2").Return(model.UserResponse{ID: "user-2", Username: "user2", Email: "user2@gmail.com", Created_At: time.Now()}, nil)
	mockCatRepo.EXPECT().IsUserCategoryExist("user-2", "category1").Return(true)

	// Mock for Create Category with invalid user id
	mockUserRepo.EXPECT().GetUser("invalid_id").Return(model.UserResponse{}, nil)

	t.Run("Create Category", func(t *testing.T){
		err := categoryService.CreateCategory("user-1", model.CategoryRequest{Name: "category1"})

		assert.NoError(t,err)
	})

	t.Run("Create Category with existing category name", func(t *testing.T){
		err := categoryService.CreateCategory("user-2", model.CategoryRequest{Name: "category1"})

		assert.ErrorIs(t, err, service.ErrCategoryAlreadyExist)
	})

	t.Run("Create Category with invalid user id", func(t *testing.T){
		err := categoryService.CreateCategory("invalid_id", model.CategoryRequest{Name: "category1"})

		assert.ErrorIs(t, err, service.ErrUserNotFound)
	})
}

func TestDeleteCategory(t *testing.T){
	ctrl, categoryService, mockCatRepo, _ := SetupCategoryService(t)
	defer ctrl.Finish()

	// Mock for Delete Category
	mockCatRepo.EXPECT().DeleteCategory("cat-1").Return(nil)

	// Mock for Delete Category with invalid id
	mockCatRepo.EXPECT().DeleteCategory("invalid_catid").Return(nil)

	t.Run("Delete Category", func(t *testing.T){
		err := categoryService.DeleteCategory("cat-1")

		assert.NoError(t,err)
	})

	t.Run("Delete Category with invalid id", func(t *testing.T){
		err := categoryService.DeleteCategory("invalid_catid")

		assert.ErrorIs(t, err, service.ErrCategoryNotFound)
	})
}