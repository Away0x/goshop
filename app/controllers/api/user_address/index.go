package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// Index 用户收货地址列表
func Index(c *context.AppContext, user *models.User, tokenStr string) error {
	addresses := user.Addresses()

	return c.RenderJSON(context.SerializeWrap("list", addresses))
}
