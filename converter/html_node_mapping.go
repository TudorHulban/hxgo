package converter

import (
	"strings"

	"github.com/TudorHulban/hxgo/helpers"
)

func mapNodeAttribute(nodeKey, nodeValue string) string {
	switch nodeKey {
	case "class":
		return helpers.Sprintf(
			`%s.AttrClass("%s")`,

			_DSL,
			nodeValue,
		)

	case "href":
		return helpers.Sprintf(
			`%s.Href("%s")`,

			_DSL,
			nodeValue,
		)

	default:
		return helpers.Sprintf(
			`%s.AttrWithValue("%s","%s")`,

			_DSL,
			nodeKey,
			nodeValue,
		)
	}
}

func mapNodesAttributes(data map[string]string) string {
	result := make([]string, 0, len(data))

	for node, value := range data {
		result = append(
			result,
			mapNodeAttribute(
				node,
				value,
			),
		)
	}

	return strings.Join(
		result,
		",\n",
	)
}
