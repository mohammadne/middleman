package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/middleman/internal/network"
)

type server struct {
	Servers *network.ServerConfig
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
	config.Servers = &network.ServerConfig{}

	// process
	envconfig.MustProcess("server", config)
	envconfig.MustProcess("server_server", config.Servers)
}

func (config *server) loadDev() {
	config.Servers = &network.ServerConfig{
		Host: "localhost",
		Port: "4040",
	}
}
