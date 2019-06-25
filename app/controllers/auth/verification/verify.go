/// 激活用户
package verification

import (
	"echo_shop/app/context"
)

func Verify(c *context.AppContext) error {
	return c.String(200, "verification.verify")
}
