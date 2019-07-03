package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"errors"
)

// Destroy 删除收货地址
func Destroy(c *context.AppContext, u *models.User) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("user_address", &address); err != nil {
		return err.HTML()
	}

	if err := models.Delete(address); err != nil {
		c.ErrorFlash(errors.New("收货地址删除失败: " + err.Error()))
		c.RedirectByName("user_addresses.index")
	}

	return c.RedirectByName("user_addresses.index")
}
