package helpers

import (
	"strings"
	"unicode"
)

func RemoveBeforeFirstLetter(word string) string {
	idx := strings.IndexFunc(word, unicode.IsLetter)
	if idx == -1 {
		return "" // No letters found
	}

	return word[idx:]
}
