package env

import (
	"flag"
)

var (
	// App
	AppName string
	Locale  string
	CharSet string

	// Svr
	SvrHost string
	SvrPort int

	// Db
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     int
)

func init() {
	// App
	flag.StringVar(&AppName, "AppName", "Applicant Lookup", "Application name")
	flag.StringVar(&Locale, "Locale", "en", "Language")
	flag.StringVar(&CharSet, "CharSet", "UTF-8", "Character set")

	// Svr
	flag.StringVar(&SvrHost, "SvrHost", "localhost", "Server host name")
	flag.IntVar(&SvrPort, "SvrPort", 8080, "Server port number")

	// Db
	flag.StringVar(&DbUser, "DbUser", "admin", "Db user")
	flag.StringVar(&DbPassword, "DbPassword", "secret", "Db password")
	flag.StringVar(&DbName, "DbName", "nmsrsLookup", "Db name")
	flag.StringVar(&DbHost, "DbHost", "localhost", "Db host name")
	flag.IntVar(&DbPort, "DbPort", 5984, "Db port number")

	// Parse flags
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
			"Host": SvrHost,
			"Port": string(SvrPort),
		},
		"Db": map[string]string{
			"User":     DbUser,
			"Password": DbPassword,
			"Name":     DbName,
			"Host":     DbHost,
			"Port":     string(DbPort),
		},
	}
} // TODO: Used for template access
