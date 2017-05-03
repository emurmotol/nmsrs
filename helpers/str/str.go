package str

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/zneyrl/nmsrs/helpers/fi"

	"time"

	"fmt"

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

func SnakeCaseToSentenceCase(s string) string {
	return UpperCaseFirstChar(strings.ToLower(strings.Join(strings.Split(s, "_"), " ")))
}

func SentenceCaseToSnakeCase(s string) string {
	return strings.ToLower(strings.Join(strings.Split(s, " "), "_"))
}

func AllCapsToSentenceCase(s string) string {
	return UpperCaseFirstChar(strings.ToLower(s))
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
		panic(err)
	}
	return string(hashedPwd)
}

func IsPasswordMatched(hashedPwd string, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}

func DateForHumans(sec int64) string {
	return time.Unix(sec, 0).Format("January 2, 2006")
}

func ToJSONString(in interface{}) string {
	json, err := json.Marshal(in)

	if err != nil {
		panic(err)
	}
	return string(json)
}

func BytesForHumans(b int64) string {
	kb := fi.KB
	mb := fi.MB
	gb := fi.GB
	tb := fi.TB

	switch {
	case b < kb:
		return fmt.Sprintf("%d B", b)
	case b < mb:
		return fmt.Sprintf("%d KB", b/kb)
	case b < gb:
		return fmt.Sprintf("%d MB", b/mb)
	case b < tb:
		return fmt.Sprintf("%d GB", b/gb)
	default:
		return fmt.Sprintf("%d GB", b/tb)
	}
}
