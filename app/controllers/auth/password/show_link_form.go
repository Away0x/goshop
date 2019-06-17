/// 展示发送重置密码链接 email 的页面
package password

import (
	"github.com/labstack/echo/v4"
)

func ShowLinkForm(c echo.Context) error {
	return c.String(200, "password.show_link_form")
}
