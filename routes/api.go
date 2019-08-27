package routes

import (
	"echo_shop/app/context"
	"echo_shop/pkg/constants"

	"github.com/labstack/echo/v4"
)

const (
	apiVersion = "v1"
)

var (
	APIPrefix = constants.RestfulAPIPrefix + "/" + apiVersion
)

func registerAPI(e *echo.Echo) {
	ee := e.Group(APIPrefix)

	context.RegisterHandler(ee.GET, "", func(c *context.AppContext) error {
		return c.RenderOKJSON(map[string]interface{}{
			"aa": 123,
		})
	}).Name = "api.root"
}
