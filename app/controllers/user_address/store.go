package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"

	"echo_shop/app/request"
)

// Store 创建或编辑收货地址
func Store(c *context.AppContext, u *models.User) error {
	req := new(request.AddressStoreForm)

	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectByName("user_addresses.create")
	}

	// 创建地址
	address := &models.UserAddress{UserID: u.ID}
	if err := models.AssignAndCreate(true, address, req); err != nil {
		c.ErrorFlash(c.WE(err, "收货地址创建失败"))
		return c.RedirectByName("user_addresses.create")
	}

	return c.RedirectByName("user_addresses.index")
}
