package controller

import (
	"net/http"
	"strconv"

	"github.com/hansandika/auth"
	"github.com/hansandika/pkg/users/dto"
	"github.com/hansandika/pkg/users/usecase"
	"github.com/labstack/echo"
)

type UserHTTPController struct {
	usecase usecase.UsecaseInterface
}

func InitControllerUser(uc usecase.UsecaseInterface) *UserHTTPController {
	return &UserHTTPController{
		usecase: uc,
	}
}

func (ct *UserHTTPController) CreateNewUser(c echo.Context) error {
	var input dto.NewUser
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	user, err := ct.usecase.CreateNewUser(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User added succesfully",
		"status":  http.StatusCreated,
		"data":    user,
	})
}

func (ct *UserHTTPController) GetUserById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id param",
			"status":  http.StatusBadRequest,
		})
	}

	idHeader, _ := strconv.Atoi(c.Request().Header.Get("X-Header-UserId"))
	err = auth.ValidateUser(idHeader, id)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusUnauthorized,
		})
	}

	user, err := ct.usecase.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User found",
		"status":  http.StatusOK,
		"data":    user,
	})
}

func (ct *UserHTTPController) LoginUser(c echo.Context) error {
	var input dto.UserCredential
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	token, err := ct.usecase.GetUserCredential(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User logged in succesfully",
		"status":  http.StatusOK,
		"token":   token,
	})
}

func (ct *UserHTTPController) GetAllUsers(c echo.Context) error {
	users, err := ct.usecase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users found",
		"status":  http.StatusOK,
		"data":    users,
	})
}

func (ct *UserHTTPController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id param",
			"status":  http.StatusBadRequest,
		})
	}

	idHeader, _ := strconv.Atoi(c.Request().Header.Get("X-Header-UserId"))
	err = auth.ValidateUser(idHeader, id)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusUnauthorized,
		})
	}

	var input dto.NewUser
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	user, err := ct.usecase.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
	}

	user, err = ct.usecase.UpdateUser(user, &input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated succesfully",
		"status":  http.StatusOK,
		"data":    user,
	})
}

func (ct *UserHTTPController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id param",
			"status":  http.StatusBadRequest,
		})
	}

	idHeader, _ := strconv.Atoi(c.Request().Header.Get("X-Header-UserId"))
	err = auth.ValidateUser(idHeader, id)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusUnauthorized,
		})
	}

	user, err := ct.usecase.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
	}

	user, err = ct.usecase.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted succesfully",
		"status":  http.StatusOK,
		"data":    user,
	})
}
