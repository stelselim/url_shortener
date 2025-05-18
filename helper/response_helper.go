package helper

import (
	"url_shortener/model"

	"github.com/labstack/echo/v4"
)

func RespondSuccess[T any](c echo.Context, status int, message string, data T) error {
	return c.JSON(status, model.BaseResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func RespondError(c echo.Context, status int, message string) error {
	return c.JSON(status, model.BaseResponse[any]{
		Success: false,
		Message: message,
	})
}
