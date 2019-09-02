package useraddress

import (
	"echo_shop/app/context"
	"echo_shop/app/models"

	"github.com/Away0x/validate"
)

type addressStoreForm struct {
	validate.Base
	Province     string
	City         string
	District     string
	Address      string
	Zip          uint
	ContactName  string `json:"contact_name" form:"contact_name"`
	ContactPhone string `json:"contact_phone" form:"contact_phone"`
}

func (*addressStoreForm) IsStrict() bool {
	return false
}

func (a *addressStoreForm) Validators() validate.Validators {
	return validate.Validators{
		"province": {
			validate.RequiredValidator(a.Province),
		},
		"city": {
			validate.RequiredValidator(a.City),
		},
		"district": {
			validate.RequiredValidator(a.District),
		},
		"address": {
			validate.RequiredValidator(a.Address),
		},
		"zip": {
			func() string {
				if a.Zip <= 0 {
					return "邮编不能为空"
				}
				return ""
			},
		},
		"contactName": {
			validate.RequiredValidator(a.ContactName),
		},
		"contactPhone": {
			validate.RequiredValidator(a.ContactPhone),
		},
	}
}

func (a *addressStoreForm) Messages() validate.Messages {
	return validate.Messages{
		"province":     {"省份不能为空"},
		"city":         {"城市不能为空"},
		"district":     {"区不能为空"},
		"address":      {"详细地址不能为空"},
		"zip":          {"邮编不能为空"},
		"contactName":  {"姓名不能为空"},
		"contactPhone": {"电话不能为空"},
	}
}

// Store 创建或编辑收货地址
func Store(c *context.AppContext, u *models.User) error {
	req := new(addressStoreForm)

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
