package network

import "fmt"

type ServerConfig struct {
	Host string
	Port string
}

func (cfg *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}
