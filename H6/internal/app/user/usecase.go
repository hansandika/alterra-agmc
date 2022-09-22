package user

import (
	"errors"
	"net/http"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/hansandika/internal/repository"
	"github.com/hansandika/pkg/constant"
	"github.com/hansandika/pkg/util"
	"github.com/hansandika/pkg/util/response"
)

type usecase struct {
	UserRepository repository.UserRepositoryInterface
}

type UsecaseInterface interface {
	GetUserById(id int) (*dto.UserResponse, *response.ErrorResponse)
	GetAllUsers() ([]*dto.UserResponse, *response.ErrorResponse)
	UpdateUser(id int, input *dto.NewUser) (*dto.UserResponse, *response.ErrorResponse)
	DeleteUser(id int) (*dto.UserResponse, *response.ErrorResponse)
}

func NewUsecase(f *factory.Factory) UsecaseInterface {
	return &usecase{
		UserRepository: f.UserRepository,
	}
}

func (u *usecase) GetUserById(id int) (*dto.UserResponse, *response.ErrorResponse) {
	var result *dto.UserResponse

	data, err := u.UserRepository.GetUserById(id)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, response.NewErrorResponse(http.StatusNotFound, errors.New("User not found"))
		}
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	result = &dto.UserResponse{
		ID:    int(data.ID),
		Name:  data.Name,
		Email: data.Email,
	}

	return result, nil
}

func (u *usecase) GetAllUsers() ([]*dto.UserResponse, *response.ErrorResponse) {
	var result []*dto.UserResponse

	data, err := u.UserRepository.GetAllUsers()
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	for _, user := range data {
		result = append(result, &dto.UserResponse{
			ID:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return result, nil
}

func (u *usecase) UpdateUser(id int, input *dto.NewUser) (*dto.UserResponse, *response.ErrorResponse) {
	var result *dto.UserResponse

	user, err := u.UserRepository.GetUserById(id)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, response.NewErrorResponse(http.StatusNotFound, errors.New("User not found"))
		}
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	if input.Name != "" {
		user.Name = input.Name
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		hashedPassword, err := util.HashPassword(input.Password)
		if err != nil {
			return result, response.NewErrorResponse(http.StatusInternalServerError, err)
		}
		user.Password = hashedPassword
	}

	data, err := u.UserRepository.UpdateUser(user)
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	result = &dto.UserResponse{
		ID:    int(data.ID),
		Name:  data.Name,
		Email: data.Email,
	}

	return result, nil
}

func (u *usecase) DeleteUser(id int) (*dto.UserResponse, *response.ErrorResponse) {
	var result *dto.UserResponse

	user, err := u.UserRepository.GetUserById(id)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, response.NewErrorResponse(http.StatusNotFound, errors.New("User not found"))
		}
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	err = u.UserRepository.DeleteUser(user)
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	result = &dto.UserResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
	return result, nil
}
