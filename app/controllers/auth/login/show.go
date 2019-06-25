package login

import (
	"echo_shop/app/context"
)

// Show 展示登录页面
func Show(c *context.AppContext) error {
	return c.RenderHTML("auth/login")
}
