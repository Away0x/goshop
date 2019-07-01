package useraaddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/viewmodel"
)

// Index 用户收货地址 list
func Index(c *context.AppContext, u *models.User) error {
	addresses := u.Addresses()

	return c.RenderHTML("user_addresses/index",
		viewmodel.Wrap("addresses", addresses))
}
