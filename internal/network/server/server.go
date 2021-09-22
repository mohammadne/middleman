package server

import (
	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/internal/storage"
	"github.com/mohammadne/middleman/pkg/logger"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

type restApi struct {
	config  *network.ServerConfig
	storage storage.Storage
	logger  logger.Logger

	// internal dependencies
	echo *echo.Echo
}

func New(cfg *network.ServerConfig, s storage.Storage, lg logger.Logger) *restApi {
	rest := &restApi{config: cfg, storage: s, logger: lg}

	rest.echo = echo.New()
	rest.echo.HideBanner = true
	rest.setupRoutes()

	return rest
}

func (rest *restApi) setupRoutes() {
	// rest.echo.POST("/metrics", echo.WrapHandler(promhttp.Handler()))

	booksGroup := rest.echo.Group("/objects")
	booksGroup.POST("", rest.post)
	booksGroup.GET("/:id", rest.get)
}

func (rest *restApi) Serve() {
	address := rest.config.Address()
	rest.logger.Info("starting server", logger.String("address", address))
	if err := rest.echo.Start(address); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
