package helpers

import "strings"

type Predicate[T any] func(T) T

var CanonicalDSL = []Predicate[string]{
	func(value string) string {
		return strings.ReplaceAll(value, "dsl.", "")
	},

	StripWhitespaceOutsideQuotes[string],
	RemoveNewlines[string],
	RemoveTrailingCommas[string],
	NormalizeParens[string],
	CollapseDotChains[string],
	StripComments[string],
	NormalizeCommaSpacing[string],
	RemoveDoubleSpaces[string],
	TrimOuterSpaces[string],
}

func ApplyPredicates[T any](value T, predicates ...Predicate[T]) T {
	result := value

	for _, predicate := range predicates {
		result = predicate(result)
	}

	return result
}

func StripWhitespaceOutsideQuotes[T ~string](value T) T {
	var b strings.Builder
	b.Grow(len(value))

	inQuotes := false

	for i := 0; i < len(value); i++ {
		ch := value[i]

		if ch == '"' {
			inQuotes = !inQuotes
			b.WriteByte(ch)

			continue
		}

		if !inQuotes && ch == ' ' {
			continue
		}

		b.WriteByte(ch)
	}

	return T(b.String())
}

func RemoveNewlines[T ~string](value T) T {
	var builder strings.Builder

	builder.Grow(len(value))

	for index := 0; index < len(value); index++ {
		ch := value[index]

		if ch == '\n' || ch == '\r' {
			continue
		}

		builder.WriteByte(ch)
	}

	return T(builder.String())
}

func RemoveTrailingCommas[T ~string](value T) T {
	var builder strings.Builder

	builder.Grow(len(value))

	lastWasComma := false

	for index := 0; index < len(value); index++ {
		ch := value[index]

		if ch == ',' {
			lastWasComma = true
			builder.WriteByte(ch)
			continue
		}

		if ch == ')' && lastWasComma {
			// remove the trailing comma by overwriting the last byte
			temp := builder.String()
			builder.Reset()
			builder.Grow(len(temp))

			for i := 0; i < len(temp)-1; i++ {
				builder.WriteByte(temp[i])
			}

			builder.WriteByte(')')
			lastWasComma = false

			continue
		}

		lastWasComma = false
		builder.WriteByte(ch)
	}

	return T(builder.String())
}

func NormalizeParens[T ~string](value T) T {
	var builder strings.Builder

	builder.Grow(len(value))

	previous := byte(0)

	for index := 0; index < len(value); index++ {
		ch := value[index]

		if ch == '(' {
			if previous == ' ' {
				// skip the space before '('
				temp := builder.String()
				builder.Reset()
				builder.Grow(len(temp))
				for i := 0; i < len(temp)-1; i++ {
					builder.WriteByte(temp[i])
				}
			}

			builder.WriteByte('(')
			previous = '('

			continue
		}

		if ch == ' ' && previous == '(' {
			// skip space after '('
			continue
		}

		if ch == ' ' && previous == ',' {
			// skip space after comma
			continue
		}

		builder.WriteByte(ch)
		previous = ch
	}

	return T(builder.String())
}

func CollapseDotChains[T ~string](value T) T {
	var builder strings.Builder

	builder.Grow(len(value))

	length := len(value)

	for index := 0; index < length; index++ {
		ch := value[index]

		if ch == '(' {
			// skip "()"
			if index+1 < length && value[index+1] == ')' {
				index++

				continue
			}
		}

		builder.WriteByte(ch)
	}

	return T(builder.String())
}

func StripComments[T ~string](value T) T {
	var builder strings.Builder
	builder.Grow(len(value))

	var inLineComment, inBlockComment bool

	length := len(value)
	index := 0

	for index < length {
		ch := value[index]

		if inLineComment {
			if ch == '\n' {
				inLineComment = false
				builder.WriteByte(ch)
			}
			index++

			continue
		}

		if inBlockComment {
			if ch == '*' && index+1 < length && value[index+1] == '/' {
				inBlockComment = false
				index = index + 2

				continue
			}
			index++

			continue
		}

		if ch == '/' && index+1 < length {
			next := value[index+1]

			if next == '/' {
				inLineComment = true
				index = index + 2

				continue
			}

			if next == '*' {
				inBlockComment = true
				index = index + 2

				continue
			}
		}

		builder.WriteByte(ch)
		index++
	}

	return T(builder.String())
}

func NormalizeCommaSpacing[T ~string](value T) T {
	var builder strings.Builder
	builder.Grow(len(value))

	previous := byte(0)

	for index := 0; index < len(value); index++ {
		ch := value[index]

		if ch == ' ' && previous == ',' {
			continue
		}

		builder.WriteByte(ch)
		previous = ch
	}

	return T(builder.String())
}

func RemoveDoubleSpaces[T ~string](value T) T {
	var builder strings.Builder
	builder.Grow(len(value))

	previous := byte(0)

	for index := 0; index < len(value); index++ {
		ch := value[index]

		if ch == ' ' && previous == ' ' {
			continue
		}

		builder.WriteByte(ch)
		previous = ch
	}

	return T(builder.String())
}

func TrimOuterSpaces[T ~string](value T) T {
	start := 0
	end := len(value)

	for start < end && value[start] == ' ' {
		start++
	}

	for end > start && value[end-1] == ' ' {
		end--
	}

	return T(value[start:end])
}
