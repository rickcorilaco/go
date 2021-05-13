package strings

import (
	"reflect"
)

func ExistsInSlice(element interface{}, slice interface{}) (exists bool) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return false
	}

	s := reflect.ValueOf(slice)

	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(element, s.Index(i).Interface()) {
			return true
		}
	}

	return false
}
