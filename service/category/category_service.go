package service

import (
	"errors"
	"log"

	"github.com/varomnrg/money-tracker/model"
	catRepo "github.com/varomnrg/money-tracker/repository/category"
	userRepo "github.com/varomnrg/money-tracker/repository/user"
	"github.com/varomnrg/money-tracker/utils"
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
	return c.categoryRepo.GetCategories()
}

func (c *CategoryService) GetUserCategories(userID string) ([]model.Category, error) {
	_, err := c.userRepo.GetUser(userID)
	

	log.Printf("GetUser called with userID: %s\n", userID)

	if err != nil {
		return []model.Category{}, ErrUserNotFound
	}
	return c.categoryRepo.GetUserCategories(userID)
}

func (c *CategoryService) GetCategory(id string) (model.Category, error) {
	category, err := c.categoryRepo.GetCategory(id)

	if err != nil {
		return model.Category{}, ErrCategoryNotFound
	}

	return category, nil
}

func (c *CategoryService) CreateCategory(userID string, category model.CategoryRequest) error {
	_, err := c.userRepo.GetUser(userID);

	if err != nil {
		return ErrUserNotFound
	}

	categoryExist := c.categoryRepo.IsUserCategoryExist(userID, category.Name)

	if categoryExist {
		return ErrCategoryAlreadyExist
	}

	id := "cat-" + utils.GenerateRandomID(10)

	newCategory := model.Category{
		ID: id,
		Name: category.Name,
		User_ID: userID,
	}

	return c.categoryRepo.CreateCategory(newCategory)
}

func (c *CategoryService) DeleteCategory(id string) error {
	_, err := c.categoryRepo.GetCategory(id)

	if err != nil {
		return ErrCategoryNotFound
	}

	return c.categoryRepo.DeleteCategory(id)
}

