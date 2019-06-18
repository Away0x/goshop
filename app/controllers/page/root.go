package page

import (
	"echo_shop/pkg/validate"

	"echo_shop/pkg/session"

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

// Root -
func Root(c echo.Context) error {
	s := session.NewFlash(c, "")
	s.Add("ggggg")
	s.Add("555")
	s.Save()
	return c.Render(200, "root.html", map[string]interface{}{})
}
