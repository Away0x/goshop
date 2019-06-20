package page

import (
	"echo_shop/pkg/validate"

	"echo_shop/pkg/flash"

	"github.com/labstack/echo/v4"
)

// UserForm -
type UserForm struct {
	validate.BaseForm
	// Name  string `validate:"required|minLen:7"`
	Email string `validate:"required|minLength:7"`
	// Age      int       `validate:"required|int|min:1|max:99"`
	// CreateAt int       `validate:"min:1"`
	// Safe     int       `validate:"-"`
	// UpdateAt time.Time `validate:"required"`
	// Code     string    `validate:"customValidator"` // 使用自定义验证器
}

// Root -
// func Root(c echo.Context) error {
// 	req := new(UserForm)

// 	if err := c.Bind(&req); err != nil {
// 		return c.Render(200, "root.html", map[string]interface{}{})
// 		// return errno.ParamsErr.RenderNoContent()
// 	}
// 	if err := c.Validate(req); err != nil {
// 		return errno.ParamsErr.RenderNoContent()
// 	}

// 	return errors.New("sss")
// }

var count = 1

// Root -
func Root(c echo.Context) error {
	if count%2 == 0 {
		flash.NewMessageFlash(c).
			AddSuccess("success1").
			AddSuccess("success2").
			AddInfo("info1").
			AddWarning("warning1").
			AddDanger("danger1").
			Save()
	}
	count++

	return c.Render(200, "root.html", map[string]interface{}{})
}
