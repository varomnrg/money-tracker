package repository

import "github.com/varomnrg/money-tracker/model"

type IUserRepository interface {
	GetUsers() ([]model.UserResponse, error)
	GetUser(id string) (model.UserResponse, error)
	CreateUser(movie model.User) error
	DeleteUser(id string) error
	UpdateUser(id string, movie model.UserRequest) error
	IsUsernameExist(username string) bool
}
