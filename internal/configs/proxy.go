package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
)

type proxy struct {
	Logger *logger.Config
	Server *network.ServerConfig
}

func Proxy(env string) *proxy {
	config := &proxy{}

	switch env {
	case "prod":
		config.loadProd()
	default:
		config.loadDev()
	}

	return config
}

func (config *proxy) loadProd() {
	config.Logger = &logger.Config{}
	config.Server = &network.ServerConfig{}

	// process
	envconfig.MustProcess("proxy", config)
	envconfig.MustProcess("proxy_logger", config.Logger)
	envconfig.MustProcess("proxy_server", config.Server)
}

func (config *proxy) loadDev() {
	config.Logger = &logger.Config{
		Development:      true,
		EnableCaller:     true,
		EnableStacktrace: false,
		Encoding:         "console",
		Level:            "warn",
	}

	config.Server = &network.ServerConfig{
		Host: "localhost", Port: "8090",
	}
}
