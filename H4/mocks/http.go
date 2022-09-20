package mocks

import (
	"io"
	"net/http/httptest"

	"github.com/go-playground/validator"
	validatorM "github.com/hansandika/middleware"
	"github.com/labstack/echo"
)

type EchoMock struct {
	E *echo.Echo
}

func (em *EchoMock) RequestMock(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	em.E.Validator = &validatorM.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	c := em.E.NewContext(req, rec)

	return c, rec
}
