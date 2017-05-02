package env

import (
	"flag"
	"time"

	"os"

	"github.com/zneyrl/nmsrs/helpers/fi"
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
	DefaultRegistrantPhoto    string
	DefaultMaxImageUploadSize int64

	//Key
	KeyPrivate string
	KeyPublic  string

	// JWT
	JWTTokenName string
	JWTExp       time.Duration

	// Template
	TemplateParentDir     string
	TemplateLayoutsDir    string
	TemplateExt           string
	TemplatePathSeparator string
)

func init() {
	// App
	flag.StringVar(&AppKey, "AppKey", "secret", "Application secret key")
	flag.StringVar(&AppName, "AppName", "NMSRS", "Application name")
	flag.StringVar(&Locale, "Locale", "en", "Language")
	flag.StringVar(&CharSet, "CharSet", "UTF-8", "Character set")

	// Svr
	flag.StringVar(&SvrEnvironment, "SvrEnvironment", "local", "Server environment")
	flag.StringVar(&SvrHost, "SvrHost", IP(), "Server host name")
	flag.IntVar(&SvrPort, "SvrPort", 80, "Server port number")
	flag.StringVar(&SvrProtocol, "SvrProtocol", "http", "Server protocol")

	// DB
	flag.StringVar(&DBUser, "DBUser", "admin", "Database user")
	flag.StringVar(&DBPassword, "DBPassword", "secret", "Database password")
	flag.StringVar(&DBName, "DBName", "nmsrs", "Database name")
	flag.StringVar(&DBHost, "DBHost", "localhost", "Database host name")
	flag.IntVar(&DBPort, "DBPort", 27017, "Database port number")
	flag.DurationVar(&DBTimeout, "DBTimeout", time.Duration(500*time.Millisecond), "Database timeout")

	// Svr
	flag.StringVar(&AdminName, "AdminName", "Administrator", "Administrator name")
	flag.StringVar(&AdminEmail, "AdminEmail", "admin@example.com", "Administrator email")
	flag.StringVar(&AdminPassword, "AdminPassword", "secret", "Administrator default password")

	// Default
	flag.StringVar(&DefaultUserPhoto, "DefaultUserPhoto", "/img/user/default.jpg", "Default user photo")
	flag.StringVar(&DefaultRegistrantPhoto, "DefaultRegistrantPhoto", "/img/registrant/default.jpg", "Default registrant photo")
	flag.Int64Var(&DefaultMaxImageUploadSize, "DefaultMaxImageUploadSize", int64(1*fi.MB), "Default max image upload size")

	// Key
	flag.StringVar(&KeyPrivate, "KeyPrivate", "keys/.ssh/app.rsa", "Key private rsa")  // openssl genrsa -out app.rsa keysize
	flag.StringVar(&KeyPublic, "KeyPublic", "keys/.ssh/app.rsa.pub", "Key public rsa") // openssl rsa -in app.rsa -pubout > app.rsa.pub

	// JWT
	flag.StringVar(&JWTTokenName, "JWTTokenName", "AccessToken", "JWT token name")
	flag.DurationVar(&JWTExp, "JWTExp", time.Hour*time.Duration(336), "JWT expiration") // TODO: Expires in 1 week

	// Template
	flag.StringVar(&TemplateParentDir, "TemplateParentDir", "views", "Template parent directory")
	flag.StringVar(&TemplateLayoutsDir, "TemplateLayoutsDir", "layouts", "Template layouts directory")
	flag.StringVar(&TemplateExt, "TemplateExt", ".gohtml", "Template file extension")
	flag.StringVar(&TemplatePathSeparator, "TemplatePathSeparator", string(os.PathSeparator), "Template path separator")

	// Parse flags
	flag.Parse()
}
