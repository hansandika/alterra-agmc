package controller

import (
	"net/http"
	"strconv"

	"github.com/hansandika/pkg/books/dto"
	"github.com/hansandika/pkg/books/usecase"
	"github.com/labstack/echo"
)

type BookHTTPController struct {
	usecase usecase.UsecaseInterface
}

func InitControllerBook(uc usecase.UsecaseInterface) *BookHTTPController {
	return &BookHTTPController{
		usecase: uc,
	}
}

func (ct *BookHTTPController) CreateNewBook(c echo.Context) error {
	var input dto.NewBook
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

	book, err := ct.usecase.CreateNewBook(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Book added succesfully",
		"status":  http.StatusCreated,
		"data":    book,
	})
}

func (ct *BookHTTPController) GetBookById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id param",
			"status":  http.StatusBadRequest,
		})
	}

	book, err := ct.usecase.GetBookById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book retrieved succesfully",
		"status":  http.StatusOK,
		"data":    book,
	})
}

func (ct *BookHTTPController) GetAllBooks(c echo.Context) error {
	books, err := ct.usecase.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Books retrieved succesfully",
		"status":  http.StatusOK,
		"data":    books,
	})
}

func (ct *BookHTTPController) UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id param",
			"status":  http.StatusBadRequest,
		})
	}

	var input dto.NewBook
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

	book, err := ct.usecase.GetBookById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
	}

	book, err = ct.usecase.UpdateBook(book, &input)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book updated succesfully",
		"status":  http.StatusOK,
		"data":    book,
	})
}

func (ct *BookHTTPController) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id param",
			"status":  http.StatusBadRequest,
		})
	}

	book, err := ct.usecase.GetBookById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
	}

	book, err = ct.usecase.DeleteBook(book)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book deleted succesfully",
		"status":  http.StatusOK,
		"data":    book,
	})
}
