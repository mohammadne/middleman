package proxy

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/internal/models"
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
	return network.Post(targetServer.Address(), body)
}

func (handler *restApi) get(c echo.Context) error {
	idStr := c.Param("id")

	body, err := handler.storage.Get(idStr)
	if err != nil {
		var hashId [16]byte
		copy(hashId[:], idStr)
		hashIdInt := utils.Md5HashToInt(hashId)

		targetServer := handler.serverConfigs[hashIdInt%len(handler.serverConfigs)]
		return network.Get(targetServer.Address())
	}

	return c.JSON(http.StatusOK, body)
}
