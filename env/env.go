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

	// DB
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     int
)

func init() {
	// App
	flag.StringVar(&AppName, "AppName", "Applicant Lookup", "Application name")
	flag.StringVar(&Locale, "Locale", "en", "Language")
	flag.StringVar(&CharSet, "CharSet", "UTF-8", "Character set")

	// Svr
	flag.StringVar(&SvrHost, "SvrHost", "localhost", "Server host name")
	flag.IntVar(&SvrPort, "SvrPort", 8080, "Server port number")

	// DB
	flag.StringVar(&DBUser, "DBUser", "admin", "DB user")
	flag.StringVar(&DBPassword, "DBPassword", "secret", "DB password")
	flag.StringVar(&DBName, "DBName", "nmsrs_lookup", "DB name")
	flag.StringVar(&DBHost, "DBHost", "localhost", "DB host name")
	flag.IntVar(&DBPort, "DBPort", 5984, "DB port number")

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
		"DB": map[string]string{
			"User":     DBUser,
			"Password": DBPassword,
			"Name":     DBName,
			"Host":     DBHost,
			"Port":     string(DBPort),
		},
	}
} // TODO: Used for template access
