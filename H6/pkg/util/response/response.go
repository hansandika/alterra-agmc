package response

import "github.com/labstack/echo"

type ErrorResponse struct {
	Code         int `json:"code"`
	ErrorMessage error
}

func NewErrorResponse(code int, err error) *ErrorResponse {
	return &ErrorResponse{
		Code:         code,
		ErrorMessage: err,
	}
}

func (e *ErrorResponse) SendErrorResponse(c echo.Context) error {
	return c.JSON(e.Code, map[string]interface{}{
		"message": e.ErrorMessage.Error(),
		"code":    e.Code,
	})
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(code int, message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (s *SuccessResponse) SendSuccessResponse(c echo.Context) error {
	return c.JSON(s.Code, map[string]interface{}{
		"message": s.Message,
		"code":    s.Code,
		"data":    s.Data,
	})
}
