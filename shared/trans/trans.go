package trans

import (
	"fmt"
)

func GetEq(f string, t string) string {
	switch t {
	case "required":
		return fmt.Sprintf("%s is required", f)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", f)
	}
	return ""
}
