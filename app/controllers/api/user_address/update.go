package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// Update 编辑收货地址
func Update(c *context.AppContext, user *models.User, tokenStr string) error {
	return c.RenderOKJSON()
}
