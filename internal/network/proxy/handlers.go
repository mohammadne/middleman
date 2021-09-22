package proxy

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
	networkPkg "github.com/mohammadne/middleman/pkg/network"
	"github.com/mohammadne/middleman/pkg/utils"
)

func (handler *restApi) post(c echo.Context) error {
	body := new(models.Body)
	if err := c.Bind(body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hashId := utils.NewMd5(body.Key)
	handler.storage.Save(strconv.FormatUint(hashId, 10), body)

	targetServer := getTargetServer(hashId, handler.serverConfigs)
	targetServerUrl := targetServer.Address() + "/objects"

	if err := networkPkg.Post(targetServerUrl, body); err != nil {
		handler.logger.Error("error retrieving file", logger.Error(err))
		return c.String(http.StatusBadRequest, "error retrieving file")
	}

	return c.JSON(http.StatusCreated, body)
}

func (handler *restApi) get(c echo.Context) error {
	idStr := c.Param("id")
	hashId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	body, err := handler.storage.Get(idStr)
	if err != nil {
		targetServer := getTargetServer(hashId, handler.serverConfigs)
		targetServerUrl := targetServer.Address() + "/objects"

		err = networkPkg.Get(targetServerUrl, body)
		if err != nil {
			handler.logger.Error("error retrieving file", logger.Error(err))
			return c.String(http.StatusBadRequest, "error retrieving file")
		}

		return c.JSON(http.StatusCreated, body)
	}

	return c.JSON(http.StatusOK, body)
}

func getTargetServer(hashId uint64, serverConfigs []network.ServerConfig) *network.ServerConfig {
	hashIdInt := int(hashId)
	if hashIdInt < 0 {
		hashIdInt = -1 * hashIdInt
	}

	reminder := hashIdInt % len(serverConfigs)
	return &serverConfigs[reminder]
}
