package proxy

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/pkg/logger"
	"github.com/mohammadne/middleman/pkg/network"
	"github.com/mohammadne/middleman/pkg/utils"
)

func (handler *restApi) post(c echo.Context) error {
	body := new(models.Body)
	if err := c.Bind(body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hashId := utils.NewMd5(body.Key)
	hashIdInt := utils.Md5HashToInt(hashId)
	handler.storage.Save(string(hashId[:]), body)

	targetServer := handler.serverConfigs[hashIdInt%len(handler.serverConfigs)]
	if err := network.Post(targetServer.Address(), body); err != nil {
		handler.logger.Error("error retrieving file", logger.Error(err))
		return c.String(http.StatusBadRequest, "error retrieving file")
	}

	return c.JSON(http.StatusCreated, body)
}

func (handler *restApi) get(c echo.Context) error {
	idStr := c.Param("id")

	body, err := handler.storage.Get(idStr)
	if err != nil {
		var hashId [16]byte
		copy(hashId[:], idStr)
		hashIdInt := utils.Md5HashToInt(hashId)

		targetServer := handler.serverConfigs[hashIdInt%len(handler.serverConfigs)]
		err = network.Get(targetServer.Address(), body)
		if err != nil {
			handler.logger.Error("error retrieving file", logger.Error(err))
			return c.String(http.StatusBadRequest, "error retrieving file")
		}

		return c.JSON(http.StatusCreated, body)
	}

	return c.JSON(http.StatusOK, body)
}
