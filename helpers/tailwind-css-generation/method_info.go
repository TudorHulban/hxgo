package tailwindcssgeneration

import (
	"reflect"
	"strings"
)

// MethodInfo represents a single DSL method and its mapped Tailwind class.
type MethodInfo struct {
	Name  string
	Class string
}

// ExtractAllMethods builds a dictionary of all exported, parameterless DSL
// methods on the given zero-value DSL object and their Tailwind class strings.
//
// It is completely agnostic to how or where the Tailwind CSS is obtained.
// It only inspects the DSL type and its methods.
//
// Contract:
//   - tailwindZero must be a zero-value instance of your DSL type (e.g. dsl.TW()).
//   - Only exported methods are considered.
//   - Only methods with signature: func (T) MethodName() T are considered,
//     where T is the DSL type (or assignable to it).
//   - The returned value is converted to string, split on spaces, and the
//     last token is taken as the Tailwind class.
func ExtractAllMethods(tailwindZero any) ([]MethodInfo, error) {
	if tailwindZero == nil {
		return nil, nil
	}

	v := reflect.ValueOf(tailwindZero)
	t := v.Type()

	// Normalize to non-pointer type for method inspection.
	// We will always call methods on a value.
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = reflect.New(t).Elem()
	}

	// DSL type T (e.g. dsl.TW)
	dslType := t

	var methods []MethodInfo

	numMethods := v.NumMethod()
	for ix := range numMethods {
		m := v.Type().Method(ix)

		// Only exported methods.
		if m.PkgPath != "" {
			continue
		}

		mv := v.Method(ix)
		mt := mv.Type()

		// Only methods with no parameters and a single return value.
		if mt.NumIn() != 0 || mt.NumOut() != 1 {
			continue
		}

		// Ensure returned type is the DSL type (or assignable).
		outType := mt.Out(0)
		if !outType.AssignableTo(dslType) {
			continue
		}

		// Call the method: T{}.MethodName()
		out := mv.Call(nil)[0]

		// Convert returned DSL value to string.
		s := dslValueToString(out)
		if s == "" {
			continue
		}

		class := extractTailwindClassToken(s)
		if class == "" {
			continue
		}

		// fmt.Println(
		// 	m.Name,
		// )

		methods = append(
			methods, MethodInfo{
				Name:  m.Name,
				Class: class,
			},
		)
	}

	return methods, nil
}

// dslValueToString converts the returned DSL value to a string.
// For a typical DSL defined as `type TW string`, this is sufficient.
// If the DSL type changes, adjust this conversion accordingly.
func dslValueToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	default:
		// Handle string-like aliases, e.g. type TW string.
		if v.CanInterface() {
			if s, ok := v.Interface().(string); ok {
				return s
			}
		}
	}

	return ""
}

// extractTailwindClassToken returns the last space-separated token from s.
// For example: "bg-red-500 hover:bg-red-600" => "hover:bg-red-600".
func extractTailwindClassToken(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	parts := strings.Fields(s)
	if len(parts) == 0 {
		return ""
	}

	return parts[len(parts)-1]
}
