package network

import "fmt"

type ServerConfig struct {
	Host string
	Port string
}

func (cfg *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

type ClientConfig struct {
	RequestsNumber   int `split_words:"true"`
	RequestsInterval int `split_words:"true"`
	ValueLength      int `split_words:"true"`
	KeyLength        int `split_words:"true"`
}
