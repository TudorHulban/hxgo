package helpers

import (
	"fmt"
	"reflect"
	"sort"
)

func MethodNamesOf(v any) []string {
	t := reflect.TypeOf(v)

	if t == nil {
		return nil
	}

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	names := make([]string, 0)

	vt := reflect.TypeOf(reflect.New(t).Elem().Interface())

	for method := range vt.Methods() {
		names = append(names, method.Name)
	}

	pt := reflect.PointerTo(t)

	for method := range pt.Methods() {
		names = append(names, method.Name)
	}

	sort.Strings(names)

	result := make([]string, 0, 1+len(names)+1)

	result = append(result,
		Sprintf("object: %s", t.Name()),
	)

	result = append(result, names...)

	result = append(
		result,
		fmt.Sprintf("methods: %d", len(names)),
	)

	return result
}
