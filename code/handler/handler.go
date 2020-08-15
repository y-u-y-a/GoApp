package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Index.")
	}
}

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
	}
}
