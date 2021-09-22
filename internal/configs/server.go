package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
)

type server struct {
	Logger  *logger.Config
	Servers []network.ServerConfig
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
	config.Servers = []network.ServerConfig{}

	// process
	envconfig.MustProcess("server", config)
	envconfig.MustProcess("server_logger", config.Logger)
	envconfig.MustProcess("server_server", config.Servers)
}

func (config *server) loadDev() {
	config.Logger = &logger.Config{
		Development:      true,
		EnableCaller:     true,
		EnableStacktrace: false,
		Encoding:         "console",
		Level:            "warn",
	}

	config.Servers = []network.ServerConfig{
		{Host: "localhost", Port: "4040"},
	}

}
