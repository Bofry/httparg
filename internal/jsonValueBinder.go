package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/util/reflectutil"
	"github.com/Bofry/structproto/valuebinder"
)

var (
	typeOfJsonRawMessage = reflect.TypeOf(json.RawMessage(nil))
)

var _ structproto.ValueBinder = new(JsonValueBinder)

type JsonValueBinder reflect.Value

func (binder JsonValueBinder) Bind(content interface{}) error {
	rv := reflect.Value(binder)
	return binder.bindJsonValue(rv, content)
}

func (binder JsonValueBinder) bindJsonValue(rv reflect.Value, content interface{}) error {
	rv = reflect.Indirect(assignZero(rv))

	// binding for the special types
	switch rv.Type() {
	case typeOfJsonRawMessage:
		// TODO create JsonRawMessageBinder !!!
		rv = reflect.Indirect(reflectutil.AssignZero(rv))
		b, err := json.Marshal(content)
		if err != nil {
			return &valuebinder.ValueBindingError{
				Value: content,
				Kind:  rv.Type().Name(),
				Err:   err}
		}
		var raw = json.RawMessage(b)
		rv.Set(reflect.ValueOf(raw))
		return nil
	}

	// perform normal binding
	scalarValueBinder := valuebinder.ScalarBinder(rv)
	err := scalarValueBinder.Bind(content)
	if err != nil {
		switch rv.Kind() {
		case reflect.Array, reflect.Slice:
			var jsonArray []interface{}
			{
				switch v := content.(type) {
				case []interface{}:
					jsonArray = v
				default:
					result, err := binder.marshalContent(content)
					if err != nil {
						return &valuebinder.ValueBindingError{
							Value: content,
							Kind:  rv.Type().Name(),
							Err:   err}
					}
					jsonArray, _ = result.([]interface{})
				}
			}

			if jsonArray == nil {
				return &valuebinder.ValueBindingError{
					Value: content,
					Kind:  rv.Type().Name(),
					Err:   err}
			}
			return binder.bindJsonArray(rv, jsonArray)
		case reflect.Struct:
			var jsonObject map[string]interface{}
			{
				switch v := content.(type) {
				case map[string]interface{}:
					jsonObject = v
				default:
					result, err := binder.marshalContent(content)
					if err != nil {
						return &valuebinder.ValueBindingError{
							Value: content,
							Kind:  rv.Type().Name(),
							Err:   err}
					}
					jsonObject, _ = result.(map[string]interface{})
				}
			}

			if jsonObject == nil {
				return &valuebinder.ValueBindingError{
					Value: content,
					Kind:  rv.Type().Name(),
					Err:   err}
			}
			return binder.bindJsonObject(rv, jsonObject)
		case reflect.Map:
			return fmt.Errorf("not implemented")
		default:
			return err
		}
	}
	return nil
}

func (binder JsonValueBinder) bindJsonArray(rv reflect.Value, content []interface{}) error {
	if len(content) > 0 {
		size := len(content)
		container := reflect.MakeSlice(rv.Type(), size, size)
		for i, elem := range content {
			err := binder.bindJsonValue(container.Index(i), elem)
			if err != nil {
				return err
			}
		}
		rv.Set(container)
	}
	return nil
}

func (binder JsonValueBinder) bindJsonObject(rv reflect.Value, content map[string]interface{}) error {
	prototype, err := structproto.Prototypify(rv,
		&structproto.StructProtoResolveOption{
			TagName: JsonTagName,
		})
	if err != nil {
		return err
	}
	return prototype.BindFields(content, BuildJsonValueBinder)
}

func (binder JsonValueBinder) marshalContent(content interface{}) (interface{}, error) {
	var (
		result interface{}
		err    error
	)

	switch buffer := content.(type) {
	case string:
		decoder := json.NewDecoder(strings.NewReader(buffer))
		decoder.UseNumber()
		err = decoder.Decode(&result)
	case []byte:
		decoder := json.NewDecoder(bytes.NewReader(buffer))
		decoder.UseNumber()
		err = decoder.Decode(&result)
	default:
		err = fmt.Errorf("cannot marshal content with type %T", content)
	}
	return result, err
}

func BuildJsonValueBinder(rv reflect.Value) structproto.ValueBinder {
	return JsonValueBinder(rv)
}
