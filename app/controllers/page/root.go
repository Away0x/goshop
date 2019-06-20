package page

import (
	"echo_shop/pkg/validate"

	"echo_shop/pkg/flash"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserForm -
type UserForm struct {
	validate.BaseForm
	Name string `validate:"required|minLen:7"`
	Age  int    `validate:"required|int|min:1|max:99"`
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
	if c.Request().Method == "POST" {
		req := new(UserForm)

		if err := c.Bind(req); err != nil {
			flash.NewDangerMessage(c, err.Error())
			return c.Redirect(http.StatusFound, "/")
		}

		if err := c.Validate(req); err != nil {
			flash.NewErrors(c, err)
			return c.Redirect(http.StatusFound, "/")
		}

		flash.NewMessageFlash(c).
			AddSuccess("success1").
			AddSuccess("success2").
			AddInfo("info1").
			AddWarning("warning1").
			AddDanger("danger1").
			Save()

		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "root.html", map[string]interface{}{})
}
