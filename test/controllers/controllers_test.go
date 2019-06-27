package controllers

import (
	"echo_shop/bootstrap"
	"echo_shop/config"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gavv/httpexpect"
)

const (
	// 默认配置文件路径
	defaultConfigFilePath = "config.yaml"
	// 配置文件格式
	configFileType = "yaml"
)

func TestControllers(t *testing.T) {
	os.Chdir("../..") // 移动工作目录
	config.InitConfig(defaultConfigFilePath, configFileType)

	go func() {
		bootstrap.Run()
	}()

	time.Sleep(time.Second * 5) // 等待 app 启动

	expect := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  config.String("APP.URL"),
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCompactPrinter(t),
		},
	})

	// 测试
	helloHandler(t, expect)
	rootHandler(t, expect)
}

func helloHandler(t *testing.T, e *httpexpect.Expect) {
	e.GET("/sd/health").
		Expect().
		Status(http.StatusOK).
		Body().
		Equal("OK")
}

func rootHandler(t *testing.T, e *httpexpect.Expect) {
	var (
		testName = "root handler"
		url      = config.String("APP.URL") + "/"
		title    = "首页"
	)

	res, err := http.Get(url)
	if err != nil {
		t.Errorf("%s http get error: %s", testName, err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Errorf("%s status code error: %d %s", testName, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		t.Errorf("%s parse doc error: %s", testName, err.Error())
	}

	titleContent := doc.Find("title").Text()
	if titleContent != title {
		t.Errorf("%s title error: %s != %s", testName, titleContent, title)
	}
}
