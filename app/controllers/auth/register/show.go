package register

import (
	"echo_shop/app/controllers"

	"github.com/labstack/echo/v4"
)

// Show 展示注册页面
func Show(c echo.Context) error {
	return controllers.Render(c, "auth/register", map[string]interface{}{
		"captcha": controllers.CreateCaptcha(),
	})
}
