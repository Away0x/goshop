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
	defer func() {
		if err := recover(); err != nil {
			data = val
		}
	}()
	// val 是 Serializer 类型
	if v, ok := val.(Serializer); ok {
		data = v.Serialize()
		return
	}

	// val 可能是 []Serializer 类型
	value := reflect.ValueOf(val)
	kind := value.Kind()

	switch kind {
	case reflect.Slice, reflect.Array:
		len := value.Len()
		l := make([]interface{}, 0)

		for i := 0; i < len; i++ {
			item := value.Index(i).Interface()
			l = append(l, getItemSerializeData(item))
		}

		data = l
	case reflect.Ptr:
		v := value.Elem().Interface()
		if reflect.TypeOf(v).Kind() != reflect.Ptr {
			data = getSerializeData(value.Elem().Interface())
		} else {
			data = val
		}
	default:
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

// Include (注意如 val 为 ptr 的话，不会调用其 Serialize 方法)
func (s Data) Include(key string, val interface{}) Data {
	data := getSerializeData(val)
	s[key] = data

	return s
}

// Wrap (注意如 val 为 ptr 的话，不会调用其 Serialize 方法)
func Wrap(key string, val interface{}) Data {
	s := Data{}
	data := getSerializeData(val)

	s[key] = data

	return s
}
