/// 激活用户
package verification

import (
	"github.com/labstack/echo/v4"
)

func Verify(c echo.Context) error {
	return c.String(200, "verification.verify")
}
