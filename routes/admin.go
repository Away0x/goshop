package routes

import (
	"echo_shop/app/context"
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/constants"

	"github.com/labstack/echo/v4"
)

func registeAdmin(e *echo.Echo) {
	ee := e.Group(constants.AdminWebPrefix)

	context.RegisterHandler(ee.GET, "", func(c *context.AppContext) error {
		return c.RenderHTML("admin", map[string]interface{}{
			"API_ROOT":   config.String("APP.URL") + "/" + APIPrefix,
			"SHARE_DATA": "{\"name\": \"wutong\"}",
		})
	})
}
