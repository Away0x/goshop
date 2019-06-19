package flash

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	itemSeparator      = "\x00"
	keyValueSeparator  = "\x23|||\x23"
	valueItemSeparator = "\x24|||\x24"
)

type flashData struct {
	Name        string // flash name
	Data        map[string][]string
	EchoContext echo.Context
}

// NewFlashData -
func NewFlashData(name string, c echo.Context) *flashData {
	return &flashData{
		Name:        name,
		Data:        make(map[string][]string),
		EchoContext: c,
	}
}

func (f *flashData) Set(key string, val []string) {
	f.Data[key] = val
}

func (f *flashData) Save() {
	f.EchoContext.Set(f.Name, f.Data)

	var val strings.Builder
	for k, v := range f.Data {
		val.WriteString(itemSeparator + k + keyValueSeparator + strings.Join(v, valueItemSeparator) + itemSeparator)
	}

	f.saveToCookie(val.String())
}

func (f *flashData) Read() map[string][]string {
	data := f.EchoContext.Get(f.Name)
	if data != nil {
		if d, ok := data.(map[string][]string); ok {
			f.removeCookie()
			return d
		}
	}

	cookie, err := f.EchoContext.Cookie(f.Name)
	if err != nil {
		return nil
	}
	v, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return nil
	}

	result := make(map[string][]string)
	items := strings.Split(v, itemSeparator) // 得到每一项 map 的 key value
	for _, item := range items {
		if len(item) > 0 {
			kv := strings.Split(item, keyValueSeparator)
			if (len(kv) == 2) && (kv[1] != "") {
				result[kv[0]] = strings.Split(kv[1], valueItemSeparator)
			}
		}
	}

	f.EchoContext.Set(f.Name, result)
	f.removeCookie()
	return result
}

func (f *flashData) setCookie(val string, maxAge int) {
	f.EchoContext.SetCookie(&http.Cookie{
		Name:     f.Name,
		Value:    val,
		MaxAge:   maxAge,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	})
}

func (f *flashData) saveToCookie(val string) {
	f.setCookie(val, 0)
}

func (f *flashData) removeCookie() {
	f.setCookie("", -1)
}
