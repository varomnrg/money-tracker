package service

import (
	"errors"

	"github.com/varomnrg/money-tracker/model"
	"github.com/varomnrg/money-tracker/repository"
	"github.com/varomnrg/money-tracker/utils"
)

type UserService struct {
	userRepo repository.IUserRepository
}

var (
	ErrIDIsNotValid         = errors.New("id is not valid")
	ErrUserNotFound         = errors.New("user cannot be found")
	ErrUserAlreadyExist     = errors.New("user already exist")
	ErrUsernameAlreadyExist = errors.New("username already exist")
)

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{
		userRepo: repository,
	}
}

func (u *UserService) GetUsers() ([]model.UserResponse, error) {
	return u.userRepo.GetUsers()
}

func (u *UserService) GetUser(id string) (model.UserResponse, error) {
	user, err := u.userRepo.GetUser(id)

	if err != nil {
		return model.UserResponse{}, ErrUserNotFound
	}
	return user, nil
}

func (u *UserService) CreateUser(userRequest model.UserRequest) error {
	UsernameExist := u.userRepo.IsUsernameExist(userRequest.Username)

	if UsernameExist {
		return ErrUsernameAlreadyExist
	}

	id := "user-" + utils.GenerateRandomID(10)
	time := utils.GetCurrentTime()

	user := model.User{
		ID:         id,
		Username:   userRequest.Username,
		Email:      userRequest.Email,
		Password:   userRequest.Password,
		Created_At: time,
	}

	return u.userRepo.CreateUser(user)
}

func (u *UserService) UpdateUser(id string, user model.UserRequest) error {
	_, err := u.userRepo.GetUser(id)

	if err != nil {
		return ErrUserNotFound
	}

	return u.userRepo.UpdateUser(id, user)
}

func (u *UserService) DeleteUser(id string) error {
	_, err := u.userRepo.GetUser(id)

	if err != nil {
		return ErrUserNotFound
	}

	return u.userRepo.DeleteUser(id)
}
