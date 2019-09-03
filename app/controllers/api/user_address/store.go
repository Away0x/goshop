package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// Store 新建收货地址
func Store(c *context.AppContext, user *models.User, tokenStr string) error {
	return c.RenderOKJSON()
}
