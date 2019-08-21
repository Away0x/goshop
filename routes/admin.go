package routes

import (
	"echo_shop/app/context"
	"echo_shop/bootstrap/config"

	"github.com/labstack/echo/v4"
)

func registeAdmin(e *echo.Echo, prefix string) {
	ee := e.Group(prefix)

	context.RegisterHandler(ee.GET, "", func(c *context.AppContext) error {
		return c.RenderHTML("admin", map[string]interface{}{
			"BDAPI": config.String("APP.URL"),
		})
	})
}
