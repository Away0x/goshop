package password

import (
	"echo_shop/app/controllers"

	"github.com/labstack/echo/v4"
)

// ShowResetForm 展示重置密码的页面
func ShowResetForm(c echo.Context) error {
	return controllers.Render(c, "auth/password/reset")
}
