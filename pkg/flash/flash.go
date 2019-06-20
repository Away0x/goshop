// Package flash : 消息闪现模块
package flash

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type (
	flashDataValue map[string][]string

	flashData struct {
		Name    string // flash name
		Data    flashDataValue
		Request *http.Request
	}
)

const (
	itemSeparator      = "\x25"
	keyValueSeparator  = "\x23"
	valueItemSeparator = "\x24"
)

// NewFlashData -
func NewFlashData(name string, r *http.Request) *flashData {
	return &flashData{
		Name:    name,
		Data:    make(flashDataValue),
		Request: r,
	}
}

func (f *flashData) Set(val flashDataValue) *flashData {
	f.Data = val
	return f
}

func (f *flashData) Save(w http.ResponseWriter) {
	var val strings.Builder
	for k, v := range f.Data {
		val.WriteString(itemSeparator + k + keyValueSeparator + strings.Join(v, valueItemSeparator) + itemSeparator)
	}

	f.saveToCookie(w, val.String())
}

func (f *flashData) Read(w http.ResponseWriter) flashDataValue {
	cookie, err := f.Request.Cookie(f.Name)
	if err != nil {
		return nil
	}

	v, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil
	}

	result := make(flashDataValue)
	items := strings.Split(string(v), itemSeparator) // 得到每一项 map 的 key value
	for _, item := range items {
		if len(item) > 0 {
			kv := strings.Split(item, keyValueSeparator)
			if (len(kv) == 2) && (kv[1] != "") {
				result[kv[0]] = strings.Split(kv[1], valueItemSeparator)
			}
		}
	}

	f.removeCookie(w)
	return result
}

func (f *flashData) setCookie(w http.ResponseWriter, val string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     f.Name,
		Value:    val,
		MaxAge:   maxAge,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	})
}

func (f *flashData) saveToCookie(w http.ResponseWriter, val string) {
	f.setCookie(w, base64.URLEncoding.EncodeToString([]byte(val)), 0)
}

func (f *flashData) removeCookie(w http.ResponseWriter) {
	f.setCookie(w, "", -1)
}
