/// 重新发送激活用户链接
package verification

import (
	"github.com/labstack/echo/v4"
)

func Resend(c echo.Context) error {
	return c.String(200, "verification.resend")
}
