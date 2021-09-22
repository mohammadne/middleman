package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/pkg/logger"
	"github.com/mohammadne/middleman/pkg/utils"
)

func (handler *restApi) post(c echo.Context) error {
	body := new(models.Body)
	if err := c.Bind(body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	filename := utils.NewMD5(body.Key)
	err := handler.storage.Save(filename, body)
	if err != nil {
		handler.logger.Error("error saving file", logger.Error(err))
		return c.String(http.StatusBadRequest, "error saving file")
	}

	return c.JSON(http.StatusCreated, body)
}

func (handler *restApi) get(c echo.Context) error {
	idStr := c.Param("id")

	body, err := handler.storage.Get(idStr)
	if err != nil {
		handler.logger.Error("error retrieving file", logger.Error(err))
		return c.String(http.StatusBadRequest, "error retrieving file")
	}

	return c.JSON(http.StatusCreated, body)
}
