package str

import (
	"log"
	"regexp"
	"strings"

	"time"

	"golang.org/x/crypto/bcrypt"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func LowerCaseFirstChar(s string) string {
	first := string([]rune(s)[0])
	return strings.ToLower(first) + strings.TrimPrefix(s, first)
} // TODO: Unused

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

func CamelCaseToSnakeCase(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}

func Bcrypt(pwd string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPwd)
}

func IsPasswordMatched(hashedPwd string, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}

func ToHumanDateFormat(sec int64) string {
	return time.Unix(sec, 0).Format("January 2, 2006")
}
