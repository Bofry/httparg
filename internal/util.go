package internal

import (
	"fmt"
	"reflect"
)

func Indirect(v interface{}) (reflect.Value, error) {
	var rv reflect.Value
	switch v.(type) {
	case reflect.Value:
		rv = v.(reflect.Value)
	default:
		rv = reflect.ValueOf(v)
	}

	if !rv.IsValid() {
		return reflect.Value{}, fmt.Errorf("specified argument 'v' is invalid")
	}

	for {
		switch rv.Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				rv = reflect.New(rv.Type().Elem())
			}
			rv = rv.Elem()
		default:
			return rv, nil
		}
	}
}
