package bootstrap

import (
	"echo_shop/routes"

	pongo2utils "echo_shop/pkg/pongo2"

	"github.com/flosch/pongo2"
	"github.com/h3poteto/pongo2echo"
	"github.com/labstack/echo"
)

// SetupEcho -
func SetupEcho(e *echo.Echo) {
	setupRender(e)
	// 项目静态文件配置
	e.Static("/public", "public")
	e.File("/favicon.ico", "public/favicon.ico")

	routes.Register(e)
}

func setupRender(e *echo.Echo) {
	pongo2utils.Setup(&pongo2utils.Config{
		PublicPath:  "public",
		MixFilePath: "public/mix-manifest.json",
	})

	render := pongo2echo.NewRenderer()
	render.AddDirectory("resources/views")
	e.Renderer = render

	pongo2.RegisterTag("static", pongo2utils.StaticTag) // 获取静态文件地址
	pongo2.RegisterTag("mix", pongo2utils.MixTag)       // 配合 laravel-mix 使用
}
