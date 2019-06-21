package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Render 渲染模板
func Render(c echo.Context, tplName string, data ...map[string]interface{}) error {
	if len(data) != 0 {
		return c.Render(http.StatusOK, tplName+".html", data[0])
	}

	return c.Render(http.StatusOK, tplName+".html", map[string]interface{}{})
}
