package echoinit

import (
	"echo_shop/config"
	"echo_shop/routes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetupRoute(e *echo.Echo) *routes.SpecialHandlers {
	// 项目静态文件配置
	e.Static("/public", config.String("APP.PUBLIC_DIR"))
	e.File("/favicon.ico", config.String("APP.PUBLIC_DIR")+"/favicon.ico")

	sh := routes.Register(e)

	// 输出路由
	printRoutes(e)

	return sh
}

func printRoutes(e *echo.Echo) {
	routes := make([]*echo.Route, 0)
	for _, item := range e.Routes() {
		if strings.HasPrefix(item.Name, "github.com") {
			continue
		}

		routes = append(routes, item)
	}

	routesStr, _ := json.MarshalIndent(struct {
		Count  int           `json:"count"`
		Routes []*echo.Route `json:"routes"`
	}{
		Count:  len(routes),
		Routes: routes,
	}, "", "  ")
	ioutil.WriteFile(config.String("APP.TEMP_DIR")+"/routes.json", routesStr, 0644)
}
