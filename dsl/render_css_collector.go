package dsl

import (
	"sort"
	"strings"
)

type cssRuleKey struct {
	selector string
	media    string
}

type smartCSSCollector struct {
	rules map[cssRuleKey]map[string]string // merged properties
	order []cssRuleKey                     // preserve insertion order (component order)
}

func newSmartCSSCollector() *smartCSSCollector {
	return &smartCSSCollector{
		rules: make(map[cssRuleKey]map[string]string),
		order: make([]cssRuleKey, 0),
	}
}

func (c *smartCSSCollector) Add(style Style) {
	key := cssRuleKey{
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

func (c *smartCSSCollector) String() string {
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
