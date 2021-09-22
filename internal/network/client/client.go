package client

import (
	"math/rand"
	"time"

	"github.com/labstack/gommon/random"
	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
	networkPkg "github.com/mohammadne/middleman/pkg/network"
)

type client struct {
	config      *network.ClientConfig
	proxyConfig *network.ServerConfig
	logger      logger.Logger
}

func New(cfg *network.ClientConfig, pc *network.ServerConfig, lg logger.Logger) *client {
	return &client{config: cfg, proxyConfig: pc, logger: lg}
}

func (client *client) Run() {
	rand.Seed(time.Now().UnixNano())

	interval := time.Millisecond * time.Duration(client.config.RequestsInterval)
	valueLength := uint8(client.config.ValueLength)
	keyLength := uint8(client.config.KeyLength)

	client.logger.Info("start clients")

	for index := 0; index < client.config.RequestsNumber; index++ {
		networkPkg.Post(
			client.proxyConfig.Address()+"/objects",
			&models.Body{
				Value: random.String(valueLength),
				Key:   random.String(keyLength),
				Cache: index%2 == 0,
			},
		)

		time.Sleep(interval)
	}

	client.logger.Info("finish clients")
}
