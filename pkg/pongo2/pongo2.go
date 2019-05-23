package pongo2

import (
	"github.com/labstack/echo/v4"
)

// Config -
type Config struct {
	PublicPath  string // 静态文件路径
	MixFilePath string // laravel-mix 生成的 mix-manifest.json 文件的路径
	Echo        *echo.Echo
}

var gpconfig *Config

// Setup -
func Setup(c *Config) {
	gpconfig = c
}

// 相对于项目 public 静态文件目录的地址
func publicPath(staticFilePath string) string {
	if string(staticFilePath[0]) == "/" {
		return gpconfig.PublicPath + staticFilePath
	}
	return gpconfig.PublicPath + "/" + staticFilePath
}

// 生成项目静态文件地址
func staticPath(staticFilePath string) string {
	return "/" + publicPath(staticFilePath)
}
