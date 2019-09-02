package auth

import (
	"echo_shop/app/auth/jwt"
	"echo_shop/app/context"
	"time"
)

func Logout(c *context.AppContext) error {
	jwt.Forget("", time.Duration(0))
	return nil
}
