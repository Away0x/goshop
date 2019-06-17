/// 展示重置密码的页面
package password

import (
	"github.com/labstack/echo/v4"
)

func ShowResetForm(c echo.Context) error {
	return c.String(200, "password.show_reset_form")
}
