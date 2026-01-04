package dsl

import (
	"text/template"
	"unsafe"
)

func needsEscape(value string) bool {
	for i := 0; i < len(value); i++ {
		c := value[i]

		if c == '<' || c == '>' || c == '&' || c == '"' || c == '\'' {
			return true
		}
	}

	return false
}

func AttrWithValue(name, value string) Node {
	escaped := value
	if needsEscape(value) {
		escaped = template.HTMLEscapeString(value) // creates allocation.
	}

	data := &struct { // intentionally move to heap.
		name  string
		value string
	}{
		name:  name,
		value: template.HTMLEscapeString(escaped),
	}

	return Node{
		fn:          renderAttrWithValue,
		data:        unsafe.Pointer(data),
		isAttribute: true,
	}
}

func renderAttrWithValue(a *Accumulator, p unsafe.Pointer) {
	d := (*struct {
		name  string
		value string
	})(p)

	a.Write5(
		" ",
		d.name,
		`="`,
		d.value,
		`"`,
	)
}

func Attr(name string) Node {
	return AttrWithValue(name, name)
}

func AttrIf(condition bool, name, value string) Node {
	if !condition {
		return Noop
	}

	return AttrWithValue(name, value)
}

func AttrWithName(name string) Node {
	return AttrWithValue(
		"name",
		name,
	)
}

func AttrCSS(css string) Node {
	return AttrWithValue(
		"style",
		css,
	)
}

func AttrClass(class string) Node {
	return AttrWithValue(
		"class",
		class,
	)
}

func AttrClassLength(class string) Node {
	if len(class) > 0 {
		return AttrWithValue(
			"class",
			class,
		)
	}

	return Noop
}

func AttrID(value string) Node {
	return AttrWithValue(
		"id",
		value,
	)
}

func AttrIDIf(condition bool, value string) Node {
	if condition {
		return AttrWithValue(
			"id",
			value,
		)
	}

	return Noop
}

// AttrIDLength returns css id Node if passed value valid.
func AttrIDLength(value string) Node {
	if len(value) > 0 {
		return AttrWithValue(
			"id",
			value,
		)
	}

	return Noop
}

func AttrType(class string) Node {
	return AttrWithValue(
		"type",
		class,
	)
}

func LabelFor(text string) Node {
	return AttrWithValue(
		"for",
		text,
	)
}

func Lang(language string) Node {
	return AttrWithValue(
		"lang",
		language,
	)
}

func Charset(charset string) Node {
	return AttrWithValue(
		"charset",
		charset,
	)
}

func Content(content string) Node {
	return AttrWithValue(
		"content",
		content,
	)
}

func Href(ref string) Node {
	return AttrWithValue(
		"href",
		ref,
	)
}

func MethodPOST() Node {
	return AttrWithValue(
		"method",
		"post",
	)
}

func Name(name string) Node {
	return AttrWithValue(
		"name",
		name,
	)
}

func OnClick(toRun string) Node {
	return AttrWithValue(
		"onclick",
		toRun,
	)
}

func Rel(relationship string) Node {
	return AttrWithValue(
		"rel",
		relationship,
	)
}

func Src(source string) Node {
	return AttrWithValue(
		"src",
		source,
	)
}

func Action(action string) Node {
	return AttrWithValue(
		"action",
		action,
	)
}

func ActionLength(action string) Node {
	if len(action) > 0 {
		return AttrWithValue(
			"action",
			action,
		)
	}

	return AttrWithValue(
		"action",
		"#",
	)
}

func Method(method string) Node {
	return AttrWithValue(
		"method",
		method,
	)
}

func AutocompleteOff() Node {
	return AttrWithValue(
		"autocomplete",
		"off",
	)
}
