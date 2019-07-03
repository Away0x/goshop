package utils

import (
	"fmt"
	"reflect"
)

// IsBlank 判断是否为类型空值
func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}

	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// BatchAssign 批量赋值
// 将 other struct 的值赋值给 target struct (struct 传递的需为 struct 指针)
func BatchAssign(force bool, target interface{}, other interface{}) error {
	targetType := reflect.TypeOf(target)
	targetKind := targetType.Kind()
	otherType := reflect.TypeOf(other)
	otherKind := otherType.Kind()

	if targetKind != reflect.Ptr || otherKind != reflect.Ptr {
		return fmt.Errorf("BatchAssign error: %s can not assign to %s", otherType, targetType)
	}

	targetElem := targetType.Elem()
	targetKind = targetElem.Kind()
	otherElem := otherType.Elem()
	otherKind = otherElem.Kind()

	if targetKind != reflect.Struct || otherKind != reflect.Struct {
		return fmt.Errorf("BatchAssign error: %s can not assign to %s", otherType, targetType)
	}

	count := otherElem.NumField()
	targetValueElem := reflect.ValueOf(target).Elem()
	otherValueElem := reflect.ValueOf(other).Elem()
	for i := 0; i < count; i++ {
		f := otherElem.Field(i)
		tf, ok := targetElem.FieldByName(f.Name)
		if !ok {
			continue
		}

		if f.Type != tf.Type {
			continue
		}

		fv := otherValueElem.FieldByName(f.Name)
		tfv := targetValueElem.FieldByName(f.Name)

		// 过滤掉类型空值
		if force && IsBlank(fv) {
			continue
		}

		if !tfv.CanSet() {
			continue
		}

		tfv.Set(fv)
	}

	return nil
}
