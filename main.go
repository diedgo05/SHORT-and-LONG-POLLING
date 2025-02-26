package main

import (
	dependenciesBus "polling/src/buses/infraestructure/dependencies"
	routesBus "polling/src/buses/infraestructure/http/routes"
	dependenciesChofer "polling/src/choferes/infraestructure/dependencies"
	routesChofer "polling/src/choferes/infraestructure/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//CHOFERES
	dependenciesChofer.Init()
	routesChofer.SetupRoutes(r)

	//BUSES
	dependenciesBus.InitBus()
	routesBus.Routes(r)

	r.Run(":8080")
}