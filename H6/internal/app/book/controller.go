package book

import (
	"errors"
	"net/http"
	"strconv"

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

func (co *controller) CreateNewBook(c echo.Context) error {
	var input dto.NewBook
	if err := c.Bind(&input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	if err := c.Validate(input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	res, err := co.usecase.CreateNewBook(&input)
	if err != nil {
		return err.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusCreated, "Create new book success", res).SendSuccessResponse(c)
}

func (co *controller) GetAllBooks(c echo.Context) error {
	res, err := co.usecase.GetAllBooks()
	if err != nil {
		return err.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Get all books success", res).SendSuccessResponse(c)
}

func (co *controller) GetBookById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id")).SendErrorResponse(c)
	}

	res, errs := co.usecase.GetBookById(id)
	if errs != nil {
		return errs.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Get book by id success", res).SendSuccessResponse(c)
}

func (co *controller) UpdateBookById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id")).SendErrorResponse(c)
	}

	var input dto.NewBook
	if err := c.Bind(&input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	if err := c.Validate(input); err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, err).SendErrorResponse(c)
	}

	res, errs := co.usecase.UpdateBook(id, &input)
	if errs != nil {
		return errs.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Update book by id success", res).SendSuccessResponse(c)
}

func (co *controller) DeleteBookById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.NewErrorResponse(http.StatusBadRequest, errors.New("Invalid parsing id")).SendErrorResponse(c)
	}

	res, errs := co.usecase.DeleteBook(id)
	if errs != nil {
		return errs.SendErrorResponse(c)
	}
	return response.NewSuccessResponse(http.StatusOK, "Delete book by id success", res).SendSuccessResponse(c)
}
