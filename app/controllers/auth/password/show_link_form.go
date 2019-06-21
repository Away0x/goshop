package password

import (
	"echo_shop/app/controllers"

	"github.com/labstack/echo/v4"
)

// ShowLinkForm 展示发送重置密码链接 email 的页面
func ShowLinkForm(c echo.Context) error {
	return controllers.Render(c, "auth/password/email")
}
