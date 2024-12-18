package tools

import (
	"errors"
	"reflect"
)

// BoolToInt64 将bool类型转换为int64类型
func BoolToInt64(b bool) int64 {
	if b {
		return 1
	}

	return 0
}

// StructToMap 结构体根据tag转map
func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	if v.Kind() != reflect.Struct {
		return nil, errors.New("in type is not struct")
	}

	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		tagVlaue := fi.Tag.Get(tagName)
		out[tagVlaue] = v.Field(i).Interface()
	}

	return out, nil
}
