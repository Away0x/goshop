package request

import (
	"github.com/Away0x/validate"
)

// AddressStoreForm 创建 user address
type AddressStoreForm struct {
	validate.Base
	Province     string
	City         string
	District     string
	Address      string
	Zip          uint
	ContactName  string `json:"contact_name" form:"contact_name"`
	ContactPhone string `json:"contact_phone" form:"contact_phone"`
}

func (*AddressStoreForm) IsStrict() bool {
	return false
}

func (a *AddressStoreForm) Validators() validate.Validators {
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

func (a *AddressStoreForm) Messages() validate.Messages {
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
