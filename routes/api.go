package routes

import (
	"github.com/labstack/echo/v4"
)

const (
	apiVersion = "v1"
)

func registerAPI(e *echo.Echo, apiPrefix string) {
	ee := e.Group(apiPrefix + "/" + apiVersion)

	ee.GET("", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"test": "test",
		})
	}).Name = "api.root"
}
