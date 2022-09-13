package controller

import (
	"net/http"
	"strconv"

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

	var input dto.NewUser
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	user, err := ct.usecase.UpdateUser(id, &input)
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

	user, err := ct.usecase.DeleteUser(id)
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
