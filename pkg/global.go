package pkg

type Config struct {
	AppName   string
	Databases map[string]struct {
		Username string
		Password string
		Database string
		Host     string
		Port     string
	}
}

// global variable
var Conf Config
