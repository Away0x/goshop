package controllers

import (
	"echo_shop/bootstrap"
	"echo_shop/config"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
)

const (
	configPath = "../../config.yaml" // 配置文件相对位置
)

func TestControllers(t *testing.T) {
	go func() {
		bootstrap.RunWithConfig(configPath, "yaml")
	}()

	time.Sleep(time.Second * 5) // 等待 app 启动

	expect := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  config.String("APP.URL"),
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCompactPrinter(t),
		},
	})

	testControllers(expect)
}

func testControllers(e *httpexpect.Expect) {
	// root controllers
	e.GET("/").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("hello")
}
