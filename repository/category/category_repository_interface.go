package repository

import "github.com/varomnrg/money-tracker/model"

type ICategoryRepository interface {
	GetCategories() ([]model.Category, error)
	GetUserCategories(userID string) ([]model.Category, error)
	GetCategory(id string) (model.Category, error)
	CreateCategory(category model.Category) error
	DeleteCategory(id string) error
	IsUserCategoryExist(userID string, categoryName string) bool
}
