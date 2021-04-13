package ui_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello there")
}
