package usecase

import (
	"petadopter/config"
	"petadopter/domain"
	"petadopter/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

func TestRegisterUser(t *testing.T) {
	repo := new(mocks.UserData)
	token := oauth2.Token{AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiUm9sZSI6InVzZXIifQ.Bt_3k60s-9Di2KH2pHE9g2m_NyZkgm0HYncVmddcG0U"}

	userInfo := domain.UserInfo{Fullname: "Bruce Wayne", Email: "brucewayne@gmail.com", Photoprofile: "batman.jpg"}

	mockData := domain.User{Username: "batman", Fullname: "Bruce Wayne", Email: "brucewayne@gmail.com", Address: "jln.cijantung",
		City: "Jakarta", Password: "polar", Phonenumber: "081212121212"}

	returnData := mockData
	returnData.ID = 1
	returnData.Role = "user"

	invalidData := domain.User{Fullname: "Bruce Wayne", Email: "brucewayne@gmail.com", Address: "jln.cijantung",
		City: "Jakarta", Password: "polar", Phonenumber: "081212121212"}

	t.Run("Success register without google", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, config.COST, nil, domain.UserInfo{})

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Success register with google", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, config.COST, &token, userInfo)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(invalidData, config.COST, nil, domain.UserInfo{})

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicated data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, config.COST, nil, domain.UserInfo{})

		assert.Equal(t, 409, status)
		repo.AssertExpectations(t)
	})

	t.Run("Generate bcrypt error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, 40, nil, domain.UserInfo{})

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		returnData.ID = 0
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.RegisterUser(mockData, config.COST, nil, domain.UserInfo{})

		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{Username: "batman", Email: "brucewayne@gmail.com", Address: "jln.cijantung", Password: "polar", PhotoProfile: "wayne.jpg"}

	returnData := mockData
	returnData.ID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("UpdateUserData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, config.COST)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 0, config.COST)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, config.COST)

		assert.Equal(t, 409, res)
		repo.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{Username: "batman", Password: "polar"}
	returnData := domain.User{ID: 1, Role: "user", Username: "batman"}
	notfound := mockData
	notfound.ID = 0
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiUm9sZSI6InVzZXIifQ.TqvhPckNECPMraeU5gaCJnK27HNnZCs8n6TszCZrrc8"

	t.Run("Succes Login", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru")
		repo.On("Login", mock.Anything, mock.Anything).Return(returnData).Once()
		userUseCase := New(repo, validator.New())
		res, status := userUseCase.Login(mockData, nil)

		assert.Equal(t, "batman", res["username"])
		assert.Equal(t, token, res["token"])
		assert.Equal(t, "user", res["role"])
		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru")
		repo.On("Login", mock.Anything, mock.Anything).Return(domain.User{}).Once()
		userUseCase := New(repo, validator.New())
		res, status := userUseCase.Login(mockData, nil)

		assert.Nil(t, res)
		assert.Equal(t, status, 404)
		repo.AssertExpectations(t)
	})

	t.Run("Wrong input", func(t *testing.T) {
		mockData.Password = ""
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8w")
		userUseCase := New(repo, validator.New())
		res, status := userUseCase.Login(mockData, nil)

		assert.Nil(t, res)
		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(200).Once()
		usecase := New(repo, validator.New())
		status := usecase.Delete(1)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(404).Once()
		usecase := New(repo, validator.New())
		status := usecase.Delete(0)

		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})

	t.Run("Internal server error", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(500).Once()
		usecase := New(repo, validator.New())
		status := usecase.Delete(0)

		assert.Equal(t, 500, status)
		repo.AssertExpectations(t)
	})
}
func TestGetUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{ID: 1, Username: "batman", Email: "brucewayne@gmail.com", Address: "jakarta", Password: "polar", PhotoProfile: "wayne.jpg"}
	t.Run("success get data", func(t *testing.T) {
		repo.On("GetProfile", mock.Anything).Return(mockData, 200).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetProfile(1)

		assert.Equal(t, 200, status)
		assert.NotNil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetProfile", mock.Anything).Return(domain.User{}, 404).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetProfile(0)

		assert.Equal(t, 404, status)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Internal server error", func(t *testing.T) {
		repo.On("GetProfile", mock.Anything).Return(domain.User{}, 500).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetProfile(0)

		assert.Equal(t, 500, status)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}
