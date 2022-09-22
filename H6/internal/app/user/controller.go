package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	jwtUtil "github.com/hansandika/internal/pkg/util"
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

func (co *controller) GetAllUsers(c echo.Context) error {
	res, errs := co.usecase.GetAllUsers()
	if errs != nil {
		return errs.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Get all users success", res).SendSuccessResponse(c)
}

func (co *controller) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id")).SendErrorResponse(c)
	}

	idHeader, err := strconv.Atoi(c.Request().Header.Get("X-Header-UserId"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id header")).SendErrorResponse(c)
	}

	err = jwtUtil.ValidateUser(idHeader, id)
	if err != nil {
		return response.NewErrorResponse(http.StatusUnauthorized, errors.New("This action is unauthorized")).SendErrorResponse(c)
	}

	res, errs := co.usecase.GetUserById(id)
	if errs != nil {
		return errs.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Get user success", res).SendSuccessResponse(c)
}

func (co *controller) UpdateUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id")).SendErrorResponse(c)
	}

	idHeader, err := strconv.Atoi(c.Request().Header.Get("X-Header-UserId"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id header")).SendErrorResponse(c)
	}

	err = jwtUtil.ValidateUser(idHeader, id)
	if err != nil {
		return response.NewErrorResponse(http.StatusUnauthorized, errors.New("This action is unauthorized")).SendErrorResponse(c)
	}

	var input dto.NewUser
	if err := c.Bind(&input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	if err := c.Validate(input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	res, errs := co.usecase.UpdateUser(id, &input)
	if errs != nil {
		return errs.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Update user success", res).SendSuccessResponse(c)
}

func (co *controller) DeleteUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id")).SendErrorResponse(c)
	}

	idHeader, err := strconv.Atoi(c.Request().Header.Get("X-Header-UserId"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id header")).SendErrorResponse(c)
	}

	err = jwtUtil.ValidateUser(idHeader, id)
	if err != nil {
		return response.NewErrorResponse(http.StatusUnauthorized, errors.New("This action is unauthorized")).SendErrorResponse(c)
	}

	res, errs := co.usecase.DeleteUser(id)
	if errs != nil {
		return errs.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Delete user success", res).SendSuccessResponse(c)
}
