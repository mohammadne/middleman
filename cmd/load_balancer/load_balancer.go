package loadbalancer

import (
	"log"

	"github.com/labstack/echo"
	"github.com/mohammadne/go_samples/load_balancer/internal/handlers"
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
