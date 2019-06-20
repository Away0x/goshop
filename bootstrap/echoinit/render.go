package echoinit

import (
	"echo_shop/config"
	"echo_shop/pkg/constants"
	pongo2utils "echo_shop/pkg/pongo2"
	"fmt"

	"echo_shop/pkg/flash"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

// SetupRender -
func SetupRender(e *echo.Echo) {
	pongo2utils.Setup(&pongo2utils.Config{
		Echo:        e,
		PublicPath:  config.String("APP.PUBLIC_DIR"),
		MixFilePath: config.String("APP.PUBLIC_DIR") + "/mix-manifest.json",
	})

	render := pongo2utils.NewRenderer()
	render.AddDirectory(config.String("APP.TEMPLATE_DIR"))
	render.UseContextProcessor(func(echoCtx echo.Context, pongoCtx pongo2.Context) {
		other := pongo2.Context{}

		// csrf
		csrf := echoCtx.Get(constants.CsrfContexntName)
		if c, ok := csrf.(string); ok && c != "" {
			other["csrf_token"] = c
			other["csrf_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, constants.CsrfValueName, c)
			other["csrf_meta"] = fmt.Sprintf(`<meta name="csrf-token" content="%s">`, c)
		}

		// method 重写
		other["delete_method_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="DELETE">`, constants.MethodOverrideFromFormKeyName)
		other["put_method_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="PUT">`, constants.MethodOverrideFromFormKeyName)
		other["patch_method_field"] = fmt.Sprintf(`<input type="hidden" name="%s" value="PATCH">`, constants.MethodOverrideFromFormKeyName)

		// flash
		messageFlash := flash.NewMessageFlash(echoCtx).Read()
		other["messages"] = messageFlash

		pongoCtx.Update(other)
	})
	e.Renderer = render

	pongo2.RegisterTag("static", pongo2utils.StaticTag) // 获取静态文件地址
	pongo2.RegisterTag("mix", pongo2utils.MixTag)       // 配合 laravel-mix 使用
	pongo2.RegisterTag("route", pongo2utils.RouteTag)   // 通过 route name 生成 route path
}
