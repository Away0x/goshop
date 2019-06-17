/// 发送重置密码链接的 email
package password

import (
	"github.com/labstack/echo/v4"
)

func Email(c echo.Context) error {
	return c.String(200, "password.email")
}
