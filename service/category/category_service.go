package service

import (
	"errors"

	"github.com/varomnrg/money-tracker/model"
	catRepo "github.com/varomnrg/money-tracker/repository/category"
	userRepo "github.com/varomnrg/money-tracker/repository/user"
)

type CategoryService struct {
	categoryRepo catRepo.ICategoryRepository
	userRepo    userRepo.IUserRepository
}

var (
	ErrUserNotFound	 		= errors.New("user cannot be found")
	ErrCategoryNotFound 	= errors.New("category cannot be found")
	ErrCategoryAlreadyExist = errors.New("category already exist")
)

func NewCategoryService(catRepo catRepo.ICategoryRepository, userRepo userRepo.IUserRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: catRepo,
		userRepo:    userRepo,
	}
}

func (c *CategoryService) GetCategories() ([]model.Category, error) {
	panic("not implemented")
}

func (c *CategoryService) GetUserCategories(userID string) ([]model.Category, error) {
	panic("not implemented")
}

func (c *CategoryService) GetCategory(id string) (model.Category, error) {
	panic("not implemented")
}

func (c *CategoryService) CreateCategory(userID string, category model.CategoryRequest) error {
	panic("not implemented")
}

func (c *CategoryService) DeleteCategory(id string) error {
	panic("not implemented")
}

