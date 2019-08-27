package routes

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/bootstrap/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/serializer"

	"github.com/labstack/echo/v4"
)

func registeAdmin(e *echo.Echo) {
	ee := e.Group(constants.AdminWebPrefix)

	context.RegisterHandler(ee.GET, "", func(c *context.AppContext) error {
		user := models.User{
			Name:  "wutong",
			Email: "wutong0910@foxmail.com",
		}

		return c.RenderHTML("admin", serializer.Data{
			"API_ROOT":   config.String("APP.URL") + APIPrefix,
			"URL_ROOT":   config.String("APP.URL") + constants.AdminWebPrefix,
			"SHARE_DATA": string(user.Serialize().ToJSONString()),
		})
	})
}
