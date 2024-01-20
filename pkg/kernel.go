package pkg

type Config struct {
	AppName   string
	Databases map[string]DatabaseConfig
	Services  map[string]ServiceConfig
}

type DatabaseConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

type ServiceConfig struct {
	Host string
	Port string
}

// global variable
