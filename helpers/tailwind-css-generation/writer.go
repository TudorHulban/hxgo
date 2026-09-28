package tailwindcssgeneration

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func Writer(w io.Writer, dictionary map[string]string, usedMethods []string) error {
	methodsSeen := make(map[string]struct{}, len(usedMethods))

	methodsInDictionary := make([]string, 0)
	methodsMissing := make([]string, 0)

	missingSet := make(map[string]struct{}, len(usedMethods))

	for _, method := range usedMethods {
		if _, exists := dictionary[method]; exists {
			if _, s := methodsSeen[method]; !s {
				methodsSeen[method] = struct{}{}

				methodsInDictionary = append(methodsInDictionary, method)
			}
		} else {
			if _, exists := missingSet[method]; !exists {
				missingSet[method] = struct{}{}

				methodsMissing = append(methodsMissing, method)
			}
		}
	}

	if len(methodsMissing) > 0 {
		sort.Strings(methodsMissing)

		return fmt.Errorf(
			"missing methods: %s",
			strings.Join(methodsMissing, ", "),
		)
	}

	sort.Strings(methodsInDictionary)

	for _, method := range methodsInDictionary {
		if _, err := io.WriteString(w, dictionary[method]+"\n"); err != nil {
			return err
		}
	}

	return nil
}
