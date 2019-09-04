package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"
)

// Store 新建收货地址
func Store(c *context.AppContext, user *models.User, tokenStr string) error {
	req := new(request.AddressStoreForm)
	if err := c.BindAndValidate(req); err != nil {
		return err
	}

	// 创建地址
	address := &models.UserAddress{UserID: user.ID}
	if err := models.AssignAndCreate(true, address, req); err != nil {
		return err
	}

	return c.RenderJSON(address)
}
