package dsl

import (
	"sort"
	"strings"
)

type styleRuleKey struct {
	selector string
	media    string
}

type registryStyles struct {
	rules map[styleRuleKey]map[string]string // merged properties
	order []styleRuleKey                     // preserve insertion order (component order)
}

func newStylesCollector() *registryStyles {
	return &registryStyles{
		rules: make(map[styleRuleKey]map[string]string),
		order: make([]styleRuleKey, 0),
	}
}

func (c *registryStyles) Add(style Style) {
	key := styleRuleKey{
		selector: style.Selector,
		media:    style.Media,
	}

	propsMap, exists := c.rules[key]
	if !exists {
		propsMap = make(map[string]string, len(style.Props))
		c.rules[key] = propsMap
		c.order = append(c.order, key)
	}

	for i := 0; i < len(style.Props); i++ {
		propsMap[style.Props[i][0]] = style.Props[i][1]
	}
}

func (c *registryStyles) String() string {
	if len(c.rules) == 0 {
		return ""
	}

	var sb strings.Builder

	for _, key := range c.order {
		props := c.rules[key]

		if key.media != "" {
			sb.WriteString("@media ")
			sb.WriteString(key.media)
			sb.WriteString(" {\n")
		}

		sb.WriteString("  ")
		sb.WriteString(key.selector)
		sb.WriteString(" {\n")

		// Optional: sort properties for stable output
		var propNames []string
		for name := range props {
			propNames = append(propNames, name)
		}
		sort.Strings(propNames)

		for _, name := range propNames {
			sb.WriteString("    ")
			sb.WriteString(name)
			sb.WriteString(": ")
			sb.WriteString(props[name])
			sb.WriteString(";\n")
		}

		sb.WriteString("  }\n")

		if key.media != "" {
			sb.WriteString("}\n")
		}
	}

	return sb.String()
}
