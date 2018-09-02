package tools

import (
	"reflect"
)

func InArray(val interface{}, array interface{}) (bool, int) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if true == reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true, i
			}
		}
	}

	return false, -1
}

func Trans(point string) string {
	return ""
}
