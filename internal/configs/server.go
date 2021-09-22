package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/middleman/pkg/logger"
)

type server struct {
	Host   string
	Ports  []string
	Logger *logger.Config
}

func Server(env string) *server {
	config := &server{}

	switch env {
	case "prod":
		config.loadProd()
	default:
		config.loadDev()
	}

	return config
}

func (config *server) loadProd() {
	config.Logger = &logger.Config{}

	// process
	envconfig.MustProcess("server", config)
	envconfig.MustProcess("server_logger", config.Logger)
}

func (config *server) loadDev() {
	config.Host = "localhost"

	config.Ports = []string{"8080", "8081", "8082", "8083"}

	config.Logger = &logger.Config{
		Development:      true,
		EnableCaller:     true,
		EnableStacktrace: false,
		Encoding:         "console",
		Level:            "warn",
	}
}
