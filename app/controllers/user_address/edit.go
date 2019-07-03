package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/serializer"
)

// Edit 编辑收货地址
func Edit(c *context.AppContext, u *models.User) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("user_address", &address); err != nil {
		return err.HTML()
	}

	return c.RenderHTML("user_addresses/create_and_edit", serializer.Data{
		"address": address.Serialize(),
	})
}
