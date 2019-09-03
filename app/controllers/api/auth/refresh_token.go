package auth

import (
	"echo_shop/app/auth/jwt"
	"echo_shop/app/context"
)

// RefreshToken 刷新 token
func RefreshToken(c *context.AppContext, tokenStr string) error {
	newTokenInfo, err := jwt.Refresh(tokenStr)
	if err != nil {
		return err
	}

	return c.RenderJSON(newTokenInfo)
}
