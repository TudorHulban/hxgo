package helpers

import (
	"fmt"
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
	for ix := 0; ix < pt.NumMethod(); ix++ {
		names = append(names, pt.Method(ix).Name)
	}

	sort.Strings(names)

	result := []string{
		fmt.Sprintf("object: %s", t.Name()),
	}

	result = append(result, names...)

	result = append(result, fmt.Sprintf("methods: %d", len(names)))

	return strings.Join(result, "\n")
}
