package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
)

// CreateAndEdit 创建或编辑收货地址
func CreateAndEdit(c *context.AppContext, u *models.User) error {
	return c.RenderHTML("user_addresses/create_and_edit")
}
