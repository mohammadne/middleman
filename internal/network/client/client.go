package client

import (
	"fmt"

	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
	networkPkg "github.com/mohammadne/middleman/pkg/network"
	"github.com/mohammadne/middleman/pkg/utils"
)

type client struct {
	logger       logger.Logger
	proxyAddress string
}

func New(lg logger.Logger, pc *network.ServerConfig) *client {
	return &client{logger: lg, proxyAddress: pc.Address()}
}

func (client *client) Get(key string) (*models.Body, error) {
	hashId := utils.NewMd5(key)
	body := new(models.Body)

	url := fmt.Sprintf("%s/%s", client.proxyAddress, string(hashId[:]))
	err := networkPkg.Get(url, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (client *client) Post(body interface{}) error {
	return networkPkg.Post(client.proxyAddress, body)
}
