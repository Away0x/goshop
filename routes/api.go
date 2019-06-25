package routes

import (
	"echo_shop/app/context"

	"github.com/labstack/echo/v4"
)

const (
	apiVersion = "v1"
)

func registerAPI(e *echo.Echo, apiPrefix string) {
	ee := e.Group(apiPrefix + "/" + apiVersion)

	registerHandler(ee.GET, "", func(c *context.AppContext) error {
		return c.JSON(200, map[string]string{
			"test": "test",
		})
	}).Name = "api.root"
}
