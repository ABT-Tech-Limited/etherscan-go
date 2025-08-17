package etherscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
)

func StructToMap(obj any) map[string]string {
	out := make(map[string]string)
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if !value.CanInterface() {
			continue
		}

		var val any
		switch value.Kind() {
		case reflect.Ptr:
			if value.IsNil() {
				val = nil
			} else {
				val = value.Elem().Interface()
			}
		default:
			val = value.Interface()
		}

		// 使用 json tag 作为 key，如果没有就用字段名
		key := field.Tag.Get("json")
		if key == "" || key == "-" {
			key = field.Name
		}
		out[key], _ = ToStringE(val)
	}

	return out
}

// ToStringE casts any value to a string type.
func ToStringE(i any) (string, error) {
	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int32:
		return strconv.FormatInt(int64(s), 10), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(s, 10), nil
	case json.Number:
		return s.String(), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	}
	return "", errors.New("unsupported type for ToStringE: " + reflect.TypeOf(i).String())
}
