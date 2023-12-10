package service

import "github.com/varomnrg/money-tracker/model"

type IUserService interface {
	GetUsers() ([]model.UserResponse, error)
	GetUser(id string) (model.UserResponse, error)
	CreateUser(movie model.UserRequest) error
	DeleteUser(id string) error
	UpdateUser(id string, movie model.UserRequest) error
}
