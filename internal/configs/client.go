package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
)

type client struct {
	Client *network.ClientConfig
	Proxy  *network.ServerConfig
	Logger *logger.Config
}

func Client(env string) *client {
	config := &client{}

	switch env {
	case "prod":
		config.loadProd()
	default:
		config.loadDev()
	}

	return config
}

func (config *client) loadProd() {
	config.Client = &network.ClientConfig{}
	config.Proxy = &network.ServerConfig{}
	config.Logger = &logger.Config{}

	envconfig.MustProcess("client_client", config.Client)
	envconfig.MustProcess("client_proxy", config.Proxy)
	envconfig.MustProcess("client_logger", config.Logger)
}

func (config *client) loadDev() {
	config.Client = &network.ClientConfig{
		RequestsNumber:   100,
		RequestsInterval: 200,
	}

	config.Proxy = &network.ServerConfig{
		Host: "localhost",
		Port: "8090",
	}

	config.Logger = &logger.Config{
		Development:      true,
		EnableCaller:     true,
		EnableStacktrace: false,
		Encoding:         "console",
		Level:            "warn",
	}
}
