package env

import (
	"flag"
	"time"

	"github.com/zneyrl/nmsrs-lookup/helpers/fi"
)

var (
	// App
	AppKey  string
	AppName string
	Locale  string
	CharSet string

	// Svr
	SvrEnvironment string
	SvrHost        string
	SvrPort        int
	SvrProtocol    string

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
	DefaultUserPhoto          string
	DefaultMaxImageUploadSize int64
)

func init() {
	// App
	flag.StringVar(&AppKey, "AppKey", "secret", "Application secret key")
	flag.StringVar(&AppName, "AppName", "Applicant Lookup", "Application name")
	flag.StringVar(&Locale, "Locale", "en", "Language")
	flag.StringVar(&CharSet, "CharSet", "UTF-8", "Character set")

	// Svr
	flag.StringVar(&SvrEnvironment, "SvrEnvironment", "local", "Server environment")
	flag.StringVar(&SvrHost, "SvrHost", "localhost", "Server host name")
	flag.IntVar(&SvrPort, "SvrPort", 8080, "Server port number")
	flag.StringVar(&SvrProtocol, "SvrProtocol", "http", "Server protocol")

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
	flag.StringVar(&DefaultUserPhoto, "DefaultUserPhoto", "/img/user/default.jpg", "Default user photo") // TODO: Convert to .jpg
	flag.Int64Var(&DefaultMaxImageUploadSize, "DefaultMaxImageUploadSize", int64(1*fi.MB), "Default max image upload size")

	// Parse flags
	flag.Parse()
}
