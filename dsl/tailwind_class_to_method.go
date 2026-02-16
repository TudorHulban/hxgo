package dsl

import "strings"

func (Tailwind) mapping() map[string]string {
	return map[string]string{
		"text-sm":     "TextSm",
		"sm:text-sm":  "TextSmSm",
		"md:text-sm":  "TextSmMd",
		"lg:text-sm":  "TextSmLg",
		"xl:text-sm":  "TextSmXl",
		"2xl:text-sm": "TextSmX2l",
	}
}

func parseTailwindClasses(classValue string) []string {
	raw := strings.Fields(classValue) // splits on any whitespace

	result := make([]string, 0, len(raw))

	for _, token := range raw {
		if token != "" {
			result = append(result, token)
		}
	}

	return result
}
