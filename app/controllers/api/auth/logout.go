package auth

import (
	"echo_shop/app/context"
	"echo_shop/app/helpers/jwt"
	"time"
)

func Logout(c *context.AppContext) error {
	jwt.Forget("", time.Duration(0))
	return nil
}
