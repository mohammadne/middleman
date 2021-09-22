package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/internal-new/network"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

type restApi struct {
	config *network.ServerConfig

	// internal dependencies
	echo *echo.Echo
}

func NewServer(cfg *network.ServerConfig) *restApi {
	rest := &restApi{config: cfg}

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

func (rest *restApi) Serve(<-chan struct{}) {
	address := fmt.Sprintf("%s:%s", rest.config.Host, rest.config.Port)
	// rest.logger.Info("starting server", logger.String("address", address))
	if err := rest.echo.Start(address); err != nil {
		// rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
