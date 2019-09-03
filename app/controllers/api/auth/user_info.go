package auth

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// UserInfo 获取用户信息
func UserInfo(c *context.AppContext, user *models.User, tokenStr string) error {
	return c.RenderJSON(user)
}
