package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/varomnrg/money-tracker/model"
	mock_repository "github.com/varomnrg/money-tracker/repository"
	"github.com/varomnrg/money-tracker/service"
	"go.uber.org/mock/gomock"
)

func SetupUserService(t *testing.T) (*gomock.Controller, *service.UserService, *mock_repository.MockIUserRepository) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repository.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	return ctrl, userService, mockRepo
}

func TestGetUsers(t *testing.T) {
	ctrl, userService, mockRepo := SetupUserService(t)
	defer ctrl.Finish()

	mockRepo.EXPECT().GetUsers().Return([]model.UserResponse{{ID: "1", Username: "user1", Email: "user1@example.com"}}, nil)

	t.Run("Get Users", func(t *testing.T) {
		users, err := userService.GetUsers()

		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, "user1", users[0].Username)
	})
}

func TestGetUser(t *testing.T) {
	ctrl, userService, mockRepo := SetupUserService(t)
	defer ctrl.Finish()

	mockRepo.EXPECT().GetUser("1").Return(model.UserResponse{ID: "1", Username: "user1", Email: "user1@example.com"}, nil)
	mockRepo.EXPECT().GetUser("invalid_id").Return(model.UserResponse{}, errors.New("user cannot be found"))

	t.Run("Get User", func(t *testing.T) {
		user, err := userService.GetUser("1")

		assert.NoError(t, err)
		assert.Equal(t, "user1", user.Username)
	})

	t.Run("Get User with invalid id", func(t *testing.T) {
		_, err := userService.GetUser("invalid_id")

		assert.ErrorIs(t, err, service.ErrUserNotFound)
	})
}

func TestCreateUser(t *testing.T) {
	ctrl, userService, mockRepo := SetupUserService(t)
	defer ctrl.Finish()

	mockRepo.EXPECT().IsUsernameExist("existing_username").Return(true)
	mockRepo.EXPECT().IsUsernameExist("new_username").Return(false)
	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)

	t.Run("Create User", func(t *testing.T) {
		err := userService.CreateUser(model.UserRequest{Username: "existing_username"})
		assert.ErrorIs(t, err, service.ErrUsernameAlreadyExist)
	})

	t.Run("Create User with new username", func(t *testing.T) {
		err := userService.CreateUser(model.UserRequest{Username: "new_username"})
		assert.NoError(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl, userService, mockRepo := SetupUserService(t)
	defer ctrl.Finish()

	mockRepo.EXPECT().GetUser("user-1").Return(model.UserResponse{ID: "user-1", Username: "user1", Email: "user1@gmail.com", Created_At: time.Now()}, nil)
	mockRepo.EXPECT().GetUser("invalid_id").Return(model.UserResponse{}, errors.New("user cannot be found"))
	mockRepo.EXPECT().UpdateUser("user-1", gomock.Any()).Return(nil)

	t.Run("Update User", func(t *testing.T) {
		err := userService.UpdateUser("user-1", model.UserRequest{Username: "user1", Email: "user1up@gmail.com", Password: "user1"})
		assert.NoError(t, err)
	})

	t.Run("Update User with invalid id", func(t *testing.T) {
		err := userService.UpdateUser("invalid_id", model.UserRequest{Username: "user1", Email: "user1up@gmail.com", Password: "user1"})
		assert.ErrorIs(t, err, service.ErrUserNotFound)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl, userService, mockRepo := SetupUserService(t)
	defer ctrl.Finish()
	
	mockRepo.EXPECT().GetUser("user-1").Return(model.UserResponse{ID: "user-1", Username: "user1", Email: "user1@gmail.com", Created_At: time.Now()}, nil)
	mockRepo.EXPECT().GetUser("invalid_id").Return(model.UserResponse{}, errors.New("user cannot be found"))
	mockRepo.EXPECT().DeleteUser("user-1").Return(nil)

	t.Run("Delete User", func(t *testing.T) {
		err := userService.DeleteUser("user-1")
		assert.NoError(t, err)
	})

	t.Run("Delete User with invalid id", func(t *testing.T) {
		err := userService.DeleteUser("invalid_id")
		assert.ErrorIs(t, err, service.ErrUserNotFound)
	})
}
