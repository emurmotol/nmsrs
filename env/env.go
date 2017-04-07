package env

import (
	"flag"
)

var (
	AppName  string
	HostName string
	Port     int
	Locale   string
	CharSet  string
)

func init() {
	flag.StringVar(&AppName, "AppName", "Applicant Lookup", "Application name")
	flag.StringVar(&HostName, "HostName", "localhost", "Host name")
	flag.IntVar(&Port, "Port", 8080, "Port number")
	flag.StringVar(&Locale, "Locale", "en", "Language")
	flag.StringVar(&CharSet, "CharSet", "UTF-8", "Character set")
	flag.Parse()
}

func Config() interface{} {
	return map[string]interface{}{
		"App": map[string]string{
			"Name":    AppName,
			"Locale":  Locale,
			"CharSet": CharSet,
		},
		"Svr": map[string]string{
			"Host": HostName,
			"Port": string(Port),
		},
	}
}
