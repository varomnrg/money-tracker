package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/varomnrg/money-tracker/handler"
	"github.com/varomnrg/money-tracker/model"
	mock_service "github.com/varomnrg/money-tracker/service"
	"go.uber.org/mock/gomock"
)


func SetupUserHandler(t *testing.T) (*gomock.Controller, *handler.UserHandler, *mock_service.MockIUserService) {
	ctrl := gomock.NewController(t)
	mockService := mock_service.NewMockIUserService(ctrl)
	userHandler := handler.NewUserHandler(mockService)

	return ctrl, userHandler, mockService
}

func TestGetUsers_Handler(t *testing.T) {
	ctrl, userHandler, mockService := SetupUserHandler(t)
	defer ctrl.Finish()

	mockService.EXPECT().GetUsers().Return(
		[]model.UserResponse{
			{ID: "user-11", Username: "user1", Email: "user1@example.com", Created_At: time.Now()},
			{ID: "user-2", Username: "user2", Email: "user2@example.com", Created_At: time.Now()},
		},
		nil,
	)


	t.Run("Get Users", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users", nil)

		recorder := httptest.NewRecorder()

		userHandler.GetUsers(recorder, req, nil)
		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")

		var users []model.UserResponse
		json.Unmarshal(recorder.Body.Bytes(), &users)

		assert.Len(t, users, 2, "Expected one user returned")
		assert.Equal(t, "user1", users[0].Username, "Expected user1 in response array at index 0")
	})
}

func TestGetUser_Handler(t *testing.T) {
	ctrl, userHandler, mockService := SetupUserHandler(t)
	defer ctrl.Finish()

	mockService.EXPECT().GetUser("1").Return(model.UserResponse{ID: "user-1", Username: "user1", Email: "user1@example.com", Created_At: time.Now()}, nil)
	mockService.EXPECT().GetUser("invalid_id").Return(model.UserResponse{}, mock_service.ErrUserNotFound)

	t.Run("Get User", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/user-1", nil)

		recorder := httptest.NewRecorder()

		userHandler.GetUser(recorder, req, []httprouter.Param{{Key: "id", Value: "1"}})

		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")

		var user model.UserResponse
		json.Unmarshal(recorder.Body.Bytes(), &user)
		
		assert.Equal(t, "user1", user.Username, "Expected user1 in response")
	})

	t.Run("Get User with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/invalid_id", nil)
		recorder := httptest.NewRecorder()

		userHandler.GetUser(recorder, req, []httprouter.Param{{Key: "id", Value: "invalid_id"}})

		assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected status Not Found")

		var errResponse model.ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &errResponse)

		assert.Equal(t, "user cannot be found", errResponse.Error, "Expected user cannot be found")
	})
}

func TestCreateUser_Handler(t *testing.T) {
	ctrl, userHandler, mockService := SetupUserHandler(t)
	defer ctrl.Finish()

	mockService.EXPECT().CreateUser(model.UserRequest{Username: "user1", Email: "user1@example.com", Password: "password"}).Return(nil)
	mockService.EXPECT().CreateUser(model.UserRequest{Username: "existingusername", Email: "exist@example.com", Password: "password"}).Return(mock_service.ErrUsernameAlreadyExist)

	t.Run("Create User", func(t *testing.T) {
		user := model.UserRequest{
			Username: "user1",
			Email:    "user1@example.com",
			Password: "password",
		}

		userBytes, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/users", bytes.NewReader(userBytes))

		recorder := httptest.NewRecorder()
		userHandler.CreateUser(recorder, req, nil)

		assert.Equal(t, http.StatusCreated, recorder.Code, "Expected status Created")
	})

	t.Run("Create User with existing username", func(t *testing.T) {
		user := model.UserRequest{
			Username: "existingusername",
			Email:    "exist@example.com",
			Password: "password",
		}

		userBytes, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/users", bytes.NewReader(userBytes))

		recorder := httptest.NewRecorder()

		userHandler.CreateUser(recorder, req, nil)

		assert.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status Bad Request")

		var errResponse model.ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &errResponse)

		assert.Equal(t, "username already exist", errResponse.Error, "Expected username already exist")
	})

}

func TestUpdateUser_Handler(t *testing.T) {
	ctrl, userHandler, mockService := SetupUserHandler(t)
	defer ctrl.Finish()

	mockService.EXPECT().UpdateUser("user-1", model.UserRequest{Username: "user1", Email: "user1@example.com", Password: "password"}).Return(nil)
	mockService.EXPECT().UpdateUser("invalid_id", model.UserRequest{Username: "user1", Email: "user1@example.com", Password: "password"}).Return(mock_service.ErrUserNotFound)

	t.Run("Update User", func(t *testing.T) {
		user := model.UserRequest{
			Username: "user1",
			Email:    "user1@example.com",
			Password: "password",
		}

		userBytes, _ := json.Marshal(user)
		req, _ := http.NewRequest("PUT", "/users/user-1", bytes.NewReader(userBytes))

		recorder := httptest.NewRecorder()
		userHandler.UpdateUser(recorder, req, []httprouter.Param{{Key: "id", Value: "user-1"}})

		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")
	})

	t.Run("Update User with invalid id", func(t *testing.T) {
		user := model.UserRequest{
			Username: "user1",
			Email:    "user1@example.com",
			Password: "password",
		}

		userBytes, _ := json.Marshal(user)
		req, _ := http.NewRequest("PUT", "/users/invalid_id", bytes.NewReader(userBytes))

		recorder := httptest.NewRecorder()
		userHandler.UpdateUser(recorder, req, []httprouter.Param{{Key: "id", Value: "invalid_id"}})

		assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected status Not Found")

		var errResponse model.ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &errResponse)

		assert.Equal(t, "user cannot be found", errResponse.Error, "Expected user cannot be found")
	})
}

func TestDeleteUser_Handler(t *testing.T) {
	ctrl, userHandler, mockService := SetupUserHandler(t)
	defer ctrl.Finish()

	mockService.EXPECT().DeleteUser("user-1").Return(nil)
	mockService.EXPECT().DeleteUser("invalid_id").Return(mock_service.ErrUserNotFound)

	t.Run("Delete User", func(t *testing.T){

		req, _ := http.NewRequest("DELETE", "/users/user-1", nil)

		recorder := httptest.NewRecorder()
		userHandler.DeleteUser(recorder, req, []httprouter.Param{{Key: "id", Value: "user-1"}})

		assert.Equal(t, http.StatusOK, recorder.Code, "Expected status OK")
	})

	t.Run("Delete User with invalid id", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/users/invalid_id", nil)

		recorder := httptest.NewRecorder()
		userHandler.DeleteUser(recorder, req, []httprouter.Param{{Key: "id", Value: "invalid_id"}})

		assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected status Not Found")

		var errResponse model.ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &errResponse)

		assert.Equal(t, "user cannot be found", errResponse.Error, "Expected user cannot be found")
	})
}