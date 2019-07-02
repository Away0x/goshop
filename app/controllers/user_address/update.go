package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/pkg/errno"

	"github.com/Away0x/validate"
)

type addressUpdateForm struct {
	addressStoreForm
}

func (*addressUpdateForm) Validators() validate.Validators {
	return validate.Validators{}
}

// Update 编辑收货地址
func Update(c *context.AppContext, u *models.User) error {
	address := new(models.UserAddress)
	if err := c.ModelByID("user_address", &address); err != nil {
		return (err.(*errno.Errno)).HTML()
	}

	req := new(addressUpdateForm)
	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectByName("user_addresses.edit")
	}

	// database.DBManager
	// 更新

	return c.RedirectByName("user_addresses.index")
}
