package login

import (
	"echo_shop/app/controllers"

	"github.com/labstack/echo/v4"
)

// Show 展示登录页面
func Show(c echo.Context) error {
	return controllers.Render(c, "auth/login")
}
