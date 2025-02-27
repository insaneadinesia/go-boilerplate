package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/apperror"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/constants"
	"github.com/labstack/echo/v4"
)

func SetupValidator() *validator.Validate {
	v := validator.New()

	// Register your custom validator here
	// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Custom_Validation_Functions

	return v
}

func Validate(c echo.Context, s interface{}) (err error) {
	if err = c.Bind(s); err != nil {
		err = apperror.New(http.StatusBadRequest, constants.CODE_INVALID_REQUEST, err)
		return
	}

	if err = c.Validate(s); err != nil {
		err = apperror.New(http.StatusBadRequest, constants.CODE_INVALID_REQUEST, err)
		return
	}

	return
}
