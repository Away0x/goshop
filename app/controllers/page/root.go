package page

import (
	"echo_shop/app/controllers"

	"github.com/labstack/echo/v4"
)

// // UserForm -
// type UserForm struct {
// 	validate.BaseForm
// 	Name string `validate:"required|minLen:7"`
// 	Age  int    `validate:"required|int|min:1|max:99"`
// }

// // Root -
// func Root(c echo.Context) error {
// 	if c.Request().Method == http.MethodPost {
// 		req := new(UserForm)

// 		if err := c.Bind(req); err != nil {
// 			flash.NewDangerMessage(c, err.Error())
// 			return c.Redirect(http.StatusFound, config.Application.Reverse("root"))
// 		}

// 		if err := c.Validate(req); err != nil {
// 			flash.NewErrors(c, err)
// 			return c.Redirect(http.StatusFound, config.Application.Reverse("root"))
// 		}

// 		flash.NewMessageFlash(c).
// 			AddSuccess("success1").
// 			AddSuccess("success2").
// 			AddInfo("info1").
// 			AddWarning("warning1").
// 			AddDanger("danger1").
// 			Save()

// 		return c.Redirect(http.StatusFound, config.Application.Reverse("root"))
// 	}

// 	return c.Render(http.StatusOK, "root.html", map[string]interface{}{})
// }

// Root -
func Root(c echo.Context) error {
	return controllers.Render(c, "root")
	// return c.String(200, "hello")
}
