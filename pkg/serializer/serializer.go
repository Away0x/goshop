package serializer

import (
	"encoding/json"
	"reflect"
)

// Data -
type Data map[string]interface{}

// Serializer response face
type Serializer interface {
	// Serialize 序列化方法
	Serialize() Data
}

func getSerializeData(val interface{}) (data interface{}) {
	value := reflect.ValueOf(val)
	kind := value.Kind()

	if kind == reflect.Slice || kind == reflect.Array {
		len := value.Len()
		l := make([]interface{}, 0)

		for i := 0; i < len; i++ {
			item := value.Index(i).Interface()
			l = append(l, getItemSerializeData(item))
		}

		data = l
	} else {
		data = val
	}

	return
}

func getItemSerializeData(val interface{}) (data interface{}) {
	switch typed := val.(type) {
	case Serializer:
		data = typed.Serialize()
	case Data:
		data = typed
	default:
		data = typed
	}

	return data
}

// ToJSON -
func (s Data) ToJSON() map[string]interface{} {
	return map[string]interface{}(s)
}

// ToJSONString -
func (s Data) ToJSONString() string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}

	return string(b)
}

// Include -
func (s Data) Include(key string, val interface{}) Data {
	data := getSerializeData(val)
	s[key] = data

	return s
}

// Wrap -
func Wrap(key string, val interface{}) Data {
	s := Data{}
	data := getSerializeData(val)

	s[key] = data

	return s
}
