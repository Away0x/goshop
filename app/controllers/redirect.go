package controllers

import (
	"echo_shop/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Redirect 重定向
func Redirect(c echo.Context, routeName string) error {
	return c.Redirect(http.StatusFound, config.Application.Reverse(routeName))
}
