package verification

import (
	"echo_shop/app/controllers"

	"github.com/labstack/echo/v4"
)

// ShowLinkForm 展示发送激活用户链接邮件的页面
func ShowLinkForm(c echo.Context) error {
	return controllers.Render(c, "auth/verification/show")
}
