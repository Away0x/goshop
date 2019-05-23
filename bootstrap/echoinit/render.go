package echoinit

import (
	pongo2utils "echo_shop/pkg/pongo2"
	"fmt"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

// SetupRender -
func SetupRender(e *echo.Echo) {
	pongo2utils.Setup(&pongo2utils.Config{
		Echo:        e,
		PublicPath:  "public",
		MixFilePath: "public/mix-manifest.json",
	})

	render := pongo2utils.NewRenderer()
	render.AddDirectory("resources/views")
	render.UseContextProcessor(func(echoCtx echo.Context, pongoCtx pongo2.Context) {
		other := pongo2.Context{}

		csrf := echoCtx.Get("csrf")
		if c, ok := csrf.(string); ok && c != "" {
			other["csrf_token"] = c
			other["csrf_input"] = fmt.Sprintf(`<input type="hidden" value="%s">`, c)
			other["csrf_meta"] = fmt.Sprintf(`<meta name="csrf-token" content="%s">`, c)
		}

		pongoCtx.Update(other)
	})
	e.Renderer = render

	pongo2.RegisterTag("static", pongo2utils.StaticTag) // 获取静态文件地址
	pongo2.RegisterTag("mix", pongo2utils.MixTag)       // 配合 laravel-mix 使用
	pongo2.RegisterTag("route", pongo2utils.RouteTag)   // 通过 route name 生成 route path
}
