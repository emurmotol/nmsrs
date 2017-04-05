package str

import (
	"regexp"
	"strings"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func LowerCaseFirstChar(s string) string {
	first := string([]rune(s)[0])
	return strings.ToLower(first) + strings.TrimPrefix(s, first)
}

func UpperCaseFirstChar(s string) string {
	first := string([]rune(s)[0])
	return strings.ToUpper(first) + strings.TrimPrefix(s, first)
}

func CamelCaseToSentenceCase(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return UpperCaseFirstChar(strings.ToLower(strings.Join(a, " ")))
}
