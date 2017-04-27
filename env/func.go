package env

import (
	"fmt"
	"net"
	"os/user"

	"github.com/zneyrl/nmsrs/helpers/lang"
)

func URL(path string) string {
	if SvrEnvironment == "production" {
		return fmt.Sprintf("%s://%s%s", SvrProtocol, SvrHost, path)
	}
	return fmt.Sprintf("%s://%s:%d%s", SvrProtocol, SvrHost, SvrPort, path)
}

func UserHomeDir() string {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func IP() string {
	ifaces, err := net.Interfaces()

	if err != nil {
		panic(err)
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()

		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()

			if ip == nil {
				continue
			}
			return ip.String()
		}
	}

	if SvrEnvironment == "production" {
		panic(lang.En["network_not_present"])
	}
	return "" // TODO: Redirect wont work if "" must be localhost
	// return "127.0.0.1" // TODO: Remote access wont work
}
