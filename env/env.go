package env

func Get() map[string]interface{} {
	return map[string]interface{}{
		"App": map[string]string{
			"Name": "Applicant Lookup",
		},
		"Svr": map[string]string{
			"Host": "localhost",
			"Port": "8080",
		},
	}
}

type App struct {
	Name string
}

type Svr struct {
	Host string
	port int
}
