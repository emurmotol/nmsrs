package helper

import (
	"fmt"

	"strings"

	"github.com/emurmotol/nmsrs/env"
)

func BaseUrl(append string) string {
	strictSlash(append)
	protocol, _ := env.Conf.String("server.protocol")
	host, _ := env.Conf.String("server.host")
	port, _ := env.Conf.Int("server.port")

	if port == 80 {
		return protocol + host + append
	}
	return protocol + fmt.Sprintf("%s:%d", host, port) + append
}

func strictSlash(str string) {
	if str != "" {
		i := strings.Index(str, "/")

		if i != 0 {
			panic(`strictSlash: param type "string" requires a slash as the first character`)
		}
	}
}

func PhotoPath(id, typ string) string {
	switch typ {
	case "User":
		return fmt.Sprintf("/users/%s/photo", id)
	case "Registrant":
		return fmt.Sprintf("/registrants/%s/photo", id)
	default:
		return fmt.Sprintf("/%s/%s/photo", strings.ToLower(typ), id)
	}
}
