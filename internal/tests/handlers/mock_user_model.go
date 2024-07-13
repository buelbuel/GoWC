package tests

import (
	models "github.com/buelbuel/gowc/models"
	"github.com/stretchr/testify/mock"
)

type MockUserModel struct {
	mock.Mock
}

func (model *MockUserModel) CreateUser(user *models.User) error {
	args := model.Called(user)
	return args.Error(0)
}

func (model *MockUserModel) GetUserByEmail(email string) (*models.User, error) {
	args := model.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (model *MockUserModel) UpdateUser(user *models.User) error {
	args := model.Called(user)
	return args.Error(0)
}

func (model *MockUserModel) DeleteUser(id string) error {
	args := model.Called(id)
	return args.Error(0)
}

func (model *MockUserModel) GetUser(id string) (*models.User, error) {
	args := model.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}
