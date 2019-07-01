package viewmodel

import (
	"reflect"
)

// Serialize -
type Serialize map[string]interface{}

// ViewModel response face
type ViewModel interface {
	// Serialize 序列化方法
	Serialize() Serialize
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
	case ViewModel:
		data = typed.Serialize()
	case Serialize:
		data = typed
	default:
		data = typed
	}

	return data
}

// ToJSON -
func (s Serialize) ToJSON() map[string]interface{} {
	return map[string]interface{}(s)
}

// Include -
func (s Serialize) Include(key string, val interface{}) Serialize {
	data := getSerializeData(val)
	s[key] = data

	return s
}

// Wrap -
func Wrap(key string, val interface{}) Serialize {
	s := Serialize{}
	data := getSerializeData(val)

	s[key] = data

	return s
}
