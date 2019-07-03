package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"errors"
)

// Update 编辑收货地址
func Update(c *context.AppContext, u *models.User) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("user_address", &address); err != nil {
		return err.HTML()
	}

	req := new(addressStoreForm)
	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectByName("user_addresses.edit", address.ID)
	}

	if err := models.AssignAndUpdate(true, address, req); err != nil {
		c.ErrorFlash(errors.New("收货地址更新失败: " + err.Error()))
		return c.RedirectByName("user_addresses.edit", address.ID)
	}

	return c.RedirectByName("user_addresses.index")
}
