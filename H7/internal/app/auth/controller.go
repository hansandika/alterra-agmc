package auth

import (
	"net/http"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/hansandika/pkg/util/response"
	"github.com/labstack/echo"
)

type controller struct {
	usecase UsecaseInterface
}

func NewController(f *factory.Factory) *controller {
	return &controller{
		usecase: NewUsecase(f),
	}
}

func (co *controller) LoginByEmailAndPassword(c echo.Context) error {
	var input dto.UserCredential
	if err := c.Bind(&input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	if err := c.Validate(input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	res, err := co.usecase.LoginByEmailAndPassword(&input)
	if err != nil {
		return err.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Login success", res).SendSuccessResponse(c)
}

func (co *controller) RegisterUserByEmailAndPassword(c echo.Context) error {
	var input dto.NewUser
	if err := c.Bind(&input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	if err := c.Validate(input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	res, err := co.usecase.RegisterUserByEmailAndPassword(&input)
	if err != nil {
		return err.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusCreated, "Register success", res).SendSuccessResponse(c)
}
