package page

import (
	"echo_shop/pkg/errno"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

type UserForm struct {
	Name     string    `validate:"required|minLen:7"`
	Email    string    `validate:"email"`
	Age      int       `validate:"required|int|min:1|max:99"`
	CreateAt int       `validate:"min:1"`
	Safe     int       `validate:"-"`
	UpdateAt time.Time `validate:"required"`
	Code     string    `validate:"customValidator"` // 使用自定义验证器
}

// Root -
func Root(c echo.Context) error {
	req := new(UserForm)
	if err := c.Bind(&req); err != nil {
		return errno.ParamsErr.SetErrors(err)
	}
	if err := c.Validate(req); err != nil {
		// return c.JSON(200, err)
		return errno.ParamsErr.SetErrors(err)
	}

	return errors.New("sss")
}
