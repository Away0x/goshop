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

	context.RegisterHandler(ee.GET, "", func(c *context.AppContext) error {
		return c.RenderOKJSON(map[string]interface{}{
			"aa": 123,
		})
	}).Name = "api.root"
}
