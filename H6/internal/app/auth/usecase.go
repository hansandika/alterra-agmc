package auth

import (
	"errors"
	"net/http"

	"github.com/hansandika/pkg/constant"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/hansandika/internal/model"
	jwtUtil "github.com/hansandika/internal/pkg/util"
	"github.com/hansandika/internal/repository"
	"github.com/hansandika/pkg/util"
	"github.com/hansandika/pkg/util/response"
)

type usecase struct {
	UserRepository repository.UserRepositoryInterface
}

type UsecaseInterface interface {
	RegisterUserByEmailAndPassword(input *dto.NewUser) (*dto.UserResponse, *response.ErrorResponse)
	LoginByEmailAndPassword(input *dto.UserCredential) (*dto.UserResponseWithToken, *response.ErrorResponse)
}

func NewUsecase(f *factory.Factory) UsecaseInterface {
	return &usecase{
		UserRepository: f.UserRepository,
	}
}

func (u *usecase) RegisterUserByEmailAndPassword(input *dto.NewUser) (*dto.UserResponse, *response.ErrorResponse) {
	var result *dto.UserResponse

	isExists, err := u.UserRepository.ValidateUserExists(input.Email)
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}
	if isExists {
		return result, response.NewErrorResponse(http.StatusBadRequest, errors.New("Email already exists"))
	}

	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	data, err := u.UserRepository.CreateNewUser(&model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	})
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

func (u *usecase) LoginByEmailAndPassword(input *dto.UserCredential) (*dto.UserResponseWithToken, *response.ErrorResponse) {
	var result *dto.UserResponseWithToken

	user, err := u.UserRepository.GetUserByEmail(input.Email)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, response.NewErrorResponse(http.StatusNotFound, errors.New("User not found"))
		}
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	if !util.CompareHashPassword(input.Password, user.Password) {
		return result, response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid email or password"))
	}

	token, err := jwtUtil.GenerateJWT(user.ID)
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, errors.New("error when genereating a new token"))
	}

	result = &dto.UserResponseWithToken{
		UserResponse: dto.UserResponse{
			ID:    int(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
		Token: token,
	}
	return result, nil
}
