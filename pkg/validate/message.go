package validate

import (
	"github.com/gookit/validate"
)

type BaseForm struct{}

func (BaseForm) Messages() map[string]string {
	return validate.MS{
		"minLength": "{field}不能小于 %d 个字符",
		"required":  "邮箱不能为空",
		"email":     "{field}格式有误",
	}
}

func (BaseForm) Translates() map[string]string {
	return validate.MS{
		"Name":  "用户名",
		"Email": "邮箱",
	}
}
