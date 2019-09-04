package context

import (
	"echo_shop/pkg/errno"
	"echo_shop/pkg/serializer"
	"net/http"
	"reflect"
)

// G -
type G = serializer.Data

// SerializeWrap -
var SerializeWrap = serializer.Wrap

// RenderHTML 渲染 html
func (a *AppContext) RenderHTML(tplName string, data ...serializer.Data) error {
	renderTypeExt := ".html"

	if len(data) != 0 {
		return a.Context.Render(http.StatusOK, tplName+renderTypeExt, data[0].ToJSON())
	}

	return a.Context.Render(http.StatusOK, tplName+renderTypeExt, map[string]interface{}{})
}

// RenderJSON render ok json
func (a *AppContext) RenderJSON(data interface{}) error {
	// data -> serializer.Serializer，会调用其 Serialize 方法，然后 ToJSON
	// data -> serializer.Data，会调用 ToJSON
	// data -> struct，会根据其 json tag 序列化

	// var d interface{}
	// switch typed := data.(type) {
	// case serializer.Serializer:
	// 	d = typed.Serialize().ToJSON()
	// case serializer.Data:
	// 	d = typed.ToJSON()
	// default:
	// 	d = data
	// }
	// return a.JSON(http.StatusOK, d)

	return a.JSON(http.StatusOK, serializer.Serialize(data))
}

// RenderOKJSON render ok json
func (a *AppContext) RenderOKJSON() error {
	return a.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}

// RenderErrorJSON render error json
func (a *AppContext) RenderErrorJSON(err *errno.Errno) error {
	resp := make(map[string]interface{})

	resp["code"] = err.Code
	resp["msg"] = err.Message

	if err.Detail != nil && err.Detail.Content != nil {
		errLen := len(err.Detail.Content)
		// detail.context 中只有一个 kv，提取出来
		if errLen == 1 {
			for _, v := range err.Detail.Content {
				resp["errors"] = getErrorContext(v)
			}
		} else {
			resp["errors"] = getErrorContext(err.Detail.Content)
		}
	}

	return a.JSON(err.HTTPCode, resp)
}

func getErrorContext(v interface{}) interface{} {
	typed := reflect.ValueOf(v).Kind()
	if typed == reflect.Func {
		return "RenderErrorJSON type error: " + typed.String()
	}

	return v
}
