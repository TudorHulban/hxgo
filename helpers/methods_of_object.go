package helpers

import (
	"reflect"
	"sort"
	"strings"
)

func MethodNamesOf(v any) string {
	t := reflect.TypeOf(v)
	if t == nil {
		return ""
	}

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	names := make([]string, 0)

	vt := reflect.TypeOf(reflect.New(t).Elem().Interface())
	for i := 0; i < vt.NumMethod(); i++ {
		names = append(names, vt.Method(i).Name)
	}

	pt := reflect.PointerTo(t)
	for i := 0; i < pt.NumMethod(); i++ {
		names = append(names, pt.Method(i).Name)
	}

	sort.Strings(names)

	return strings.Join(names, "\n")
}
