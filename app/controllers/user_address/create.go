package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// Create 创建收货地址
func Create(c *context.AppContext, u *models.User) error {
	return c.RenderHTML("user_addresses/create_and_edit")
}
