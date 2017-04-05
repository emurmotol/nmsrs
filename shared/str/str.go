package str

import "strings"

func LowerCaseFirstChar(s string) string {
	first := string([]rune(s)[0])
	return strings.ToLower(first) + strings.TrimPrefix(s, first)
}
