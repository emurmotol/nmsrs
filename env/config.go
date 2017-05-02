package env

import (
	"strconv"
)

func Config() interface{} {
	return map[string]interface{}{
		"App": map[string]string{
			"Key":     AppKey,
			"Name":    AppName,
			"Locale":  Locale,
			"CharSet": CharSet,
		},
		"Svr": map[string]string{
			"Environment": SvrEnvironment,
			"Host":        SvrHost,
			"Port":        string(SvrPort),
			"Protocol":    SvrProtocol,
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
			"UserPhoto":          DefaultUserPhoto,
			"RegistrantPhoto":    DefaultRegistrantPhoto,
			"MaxImageUploadSize": strconv.FormatInt(DefaultMaxImageUploadSize, 10),
		},
		"Key": map[string]string{
			"Private": KeyPrivate,
			"Public":  KeyPublic,
		},
		"JWT": map[string]string{
			"TokenName": JWTTokenName,
			"Exp":       string(JWTExp),
		},
		"Template": map[string]string{
			"ParentDir":     TemplateParentDir,
			"LayoutsDir":    TemplateLayoutsDir,
			"Ext":           TemplateExt,
			"PathSeparator": TemplatePathSeparator,
		},
	}
} // TODO: Used for template access
