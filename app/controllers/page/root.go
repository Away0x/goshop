package page

import (
	"echo_shop/app/context"
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
// 			return c.RouteRedirect(http.StatusFound, config.Application.Reverse("root"))
// 		}

// 		if err := c.Validate(req); err != nil {
// 			flash.NewErrors(c, err)
// 			return c.RouteRedirect(http.StatusFound, config.Application.Reverse("root"))
// 		}

// 		flash.NewMessageFlash(c).
// 			AddSuccess("success1").
// 			AddSuccess("success2").
// 			AddInfo("info1").
// 			AddWarning("warning1").
// 			AddDanger("danger1").
// 			Save()

// 		return c.RouteRedirect(http.StatusFound, config.Application.Reverse("root"))
// 	}

// 	return c.RenderHTML(http.StatusOK, "root.html", map[string]interface{}{})
// }

// Root -
func Root(c *context.AppContext) error {
	return c.RenderHTML("root")
	// return c.String(200, "hello")
}
