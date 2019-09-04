package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"
)

// Update 编辑收货地址
func Update(c *context.AppContext, user *models.User, tokenStr string) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("id", &address); err != nil {
		return err
	}

	req := new(request.AddressStoreForm)
	if err := c.BindAndValidate(req); err != nil {
		return err
	}

	if err := models.AssignAndUpdate(true, address, req); err != nil {
		return err
	}

	return c.RenderJSON(address)
}
