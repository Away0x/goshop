/// 重置密码
package password

import (
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	return c.String(200, "password.update")
}
