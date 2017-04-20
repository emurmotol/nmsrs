package env

import (
	"fmt"
	"log"
	"os/user"
)

func URL(path string) string {
	return fmt.Sprintf("%s://%s:%d%s", SvrProtocol, SvrHost, SvrPort, path)
}

func UserHomeDir() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
