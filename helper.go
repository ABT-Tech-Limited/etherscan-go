package etherscan

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
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

		if val == nil { // keep existing skip for nil pointers/interfaces
			continue
		}

		// Parse json tag: name[,option1,option2,...]
		rawTag := field.Tag.Get("json")
		if rawTag == "-" { // explicit ignore
			continue
		}
		var key string
		omitEmpty := false
		if rawTag != "" {
			parts := strings.Split(rawTag, ",")
			namePart := parts[0]
			if namePart == "-" { // already handled, but double-check
				continue
			}
			if namePart != "" {
				key = namePart
			}
			for _, opt := range parts[1:] {
				if opt == "omitempty" {
					omitEmpty = true
				}
			}
		}
		if key == "" { // fallback to field name
			key = field.Name
		}

		if omitEmpty && isEmpty(val) {
			continue
		}

		out[key], _ = ToStringE(val)
	}

	return out
}

// isEmpty determines whether a value is the zero value according to encoding/json's omitempty rules.
func isEmpty(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return rv.Len() == 0
	case reflect.Interface, reflect.Pointer, reflect.Chan, reflect.Func:
		return rv.IsNil()
	case reflect.Struct:
		// Compare with zero value of the struct
		zero := reflect.Zero(rv.Type()).Interface()
		return reflect.DeepEqual(v, zero)
	}
	return false
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
	// Fallback: for slices, arrays, maps, structs, pointers -> JSON marshal.
	rv := reflect.ValueOf(i)
	kind := rv.Kind()
	if kind == reflect.Slice || kind == reflect.Array || kind == reflect.Map || kind == reflect.Struct || kind == reflect.Pointer {
		b, err := json.Marshal(i)
		if err == nil {
			return string(b), nil
		}
		return "", err
	}
	return "", errors.New("unsupported type for ToStringE: " + reflect.TypeOf(i).String())
}

func CopyMap(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
