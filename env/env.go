package env

import (
	"flag"
	"time"
)

var (
	// App
	AppKey  string
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
	DBTimeout  time.Duration

	//Admin
	AdminName     string
	AdminEmail    string
	AdminPassword string

	// Default
	DefaultUserPhoto string
)

func init() {
	// App
	flag.StringVar(&AppKey, "AppKey", "secret", "Application secret key")
	flag.StringVar(&AppName, "AppName", "Applicant Lookup", "Application name")
	flag.StringVar(&Locale, "Locale", "en", "Language")
	flag.StringVar(&CharSet, "CharSet", "UTF-8", "Character set")

	// Svr
	flag.StringVar(&SvrHost, "SvrHost", "localhost", "Server host name")
	flag.IntVar(&SvrPort, "SvrPort", 8080, "Server port number")

	// DB
	flag.StringVar(&DBUser, "DBUser", "admin", "Database user")
	flag.StringVar(&DBPassword, "DBPassword", "secret", "Database password")
	flag.StringVar(&DBName, "DBName", "nmsrs_lookup", "Database name")
	flag.StringVar(&DBHost, "DBHost", "localhost", "Database host name")
	flag.IntVar(&DBPort, "DBPort", 27017, "Database port number")
	flag.DurationVar(&DBTimeout, "DBTimeout", time.Duration(500*time.Millisecond), "Database timeout")

	// Svr
	flag.StringVar(&AdminName, "AdminName", "Administrator", "Administrator name")
	flag.StringVar(&AdminEmail, "AdminEmail", "admin@example.com", "Administrator email")
	flag.StringVar(&AdminPassword, "AdminPassword", "secret", "Administrator default password")

	// Default
	flag.StringVar(&DefaultUserPhoto, "DefaultUserPhoto", "/img/user/default.png", "Default user photo")

	// Parse flags
	flag.Parse()
}

func Config() interface{} {
	return map[string]interface{}{
		"App": map[string]string{
			"Key":     AppKey,
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
			"Timeout":  string(DBTimeout),
		},
		"Admin": map[string]string{
			"Name":     AdminName,
			"Email":    AdminEmail,
			"Password": AdminPassword,
		},
		"Default": map[string]string{
			"UserPhoto": DefaultUserPhoto,
		},
	}
} // TODO: Used for template access
