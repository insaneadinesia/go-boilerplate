package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DefaultResponse struct {
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
}

func Success(c echo.Context, data any) error {
	resp := DefaultResponse{
		Message: "Request Successfully Processed",
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func Default(c echo.Context, status int, data any) error {
	resp := DefaultResponse{
		Message: "Request Successfully Processed",
		Data:    data,
	}

	if status >= 400 {
		resp.Message = "Request Failed to Processed"
	}

	return c.JSON(http.StatusOK, resp)
}
