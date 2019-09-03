package auth

import (
	"echo_shop/app/auth/jwt"
	"echo_shop/app/context"
	"time"
)

// Logout 登出
func Logout(c *context.AppContext, tokenStr string) error {
	jwt.Forget(tokenStr, time.Duration(0))
	return c.RenderOKJSON()
}
