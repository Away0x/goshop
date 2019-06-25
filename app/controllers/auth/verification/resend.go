/// 重新发送激活用户链接
package verification

import (
	"echo_shop/app/context"
)

func Resend(c *context.AppContext) error {
	return c.String(200, "verification.resend")
}
