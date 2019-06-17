/// 展示发送激活用户链接邮件的页面
package verification

import (
	"github.com/labstack/echo/v4"
)

func ShowLinkForm(c echo.Context) error {
	return c.String(200, "verification.show_link_form")
}
