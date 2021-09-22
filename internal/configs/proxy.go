package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
)

type proxy struct {
	Logger      *logger.Config
	Proxy       *network.ServerConfig
	ServerHost  string   `split_words:"true"`
	ServerPorts []string `split_words:"true"`
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
	config.Proxy = &network.ServerConfig{}

	// process
	envconfig.MustProcess("proxy_logger", config.Logger)
	envconfig.MustProcess("proxy_proxy", config.Proxy)
}

func (config *proxy) loadDev() {
	config.Logger = &logger.Config{
		Development:      true,
		EnableCaller:     true,
		EnableStacktrace: false,
		Encoding:         "console",
		Level:            "warn",
	}

	config.Proxy = &network.ServerConfig{
		Host: "127.0.0.1", Port: "8090",
	}

	config.ServerHost = "127.0.0.1"

	config.ServerPorts = []string{"8080", "8081", "8082", "8083", "8084"}
}
