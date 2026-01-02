package dsl

import (
	"text/template"
)

func AttrWithValue(name, value string) Node {
	return func() NodeOutput {
		escaped := template.HTMLEscapeString(value)
		// space + name + =" + value + "
		size := 1 + len(name) + 2 + len(escaped) + 1
		html := make([]byte, 0, size)

		html = append(html, ' ')
		html = append(html, name...)
		html = append(html, '=', '"')
		html = append(html, escaped...)
		html = append(html, '"')

		return NodeOutput{
			IsAttr: true,
			HTML:   html,
			Styles: nil,
		}
	}
}

func Attr(name string) Node {
	return AttrWithValue(name, name)
}

func AttrIf(condition bool, name, value string) Node {
	if !condition {
		return noop
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

	return noop
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

	return noop
}

// AttrIDLength returns css id Node if passed value valid.
func AttrIDLength(value string) Node {
	if len(value) > 0 {
		return AttrWithValue(
			"id",
			value,
		)
	}

	return noop
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
