package loadbalancer

import (
	"log"

	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/internal/handlers"
)

func Setup(path string, servers []string) {
	e := echo.New()
	e.HideBanner = true
	handlers.SetupLoadBalancerRoutes(servers, e)

	err := e.Start(path)
	if err != nil {
		log.Fatal(err)
	}
}
