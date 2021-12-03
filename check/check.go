package check

import (
	"errors"
	"reflect"
)

// Do 校验必填项 使用Tag：bind标识 msg描述错误信息
func Do(obj interface{}) (err error) {
	v := reflect.Indirect(reflect.ValueOf(obj))
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag
		if _, ok := tag.Lookup("bind"); ok {
			f := v.Field(i)
			if !validate(f) {
				return errors.New(tag.Get("msg"))
			}
		}
	}
	return nil
}

func validate(field reflect.Value) bool {
	switch field.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Int() > 0
	case reflect.String:
		return len(field.String()) > 0
	}
	return true
}
