package helper

import (
	"fmt"

	"strings"

	"github.com/emurmotol/nmsrs.v4/env"
)

func ApiBasePath(append string) string {
	strictSlash(append)
	api, _ := env.Conf.String("server.basepath.api")
	return api + append
}

func BaseURL(append string) string {
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

func PhotoPath(id int64, typ string) string {
	switch typ {
	case "User":
		return fmt.Sprintf("/users/%d/photo", id)
	case "Registrant":
		return fmt.Sprintf("/registrants/%d/photo", id)
	default:
		return fmt.Sprintf("/%s/%d/photo", strings.ToLower(typ), id)
	}
}
