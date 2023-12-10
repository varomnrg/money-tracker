package service

import "github.com/varomnrg/money-tracker/model"

type ICategoryService interface {
	GetCategories() ([]model.Category, error)
	GetUserCategories(userID string) ([]model.Category, error)
	GetCategory(id string) (model.Category, error)
	CreateCategory(userID string, category model.CategoryRequest) error
	DeleteCategory(id string) error
}
