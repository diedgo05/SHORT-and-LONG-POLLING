package routes

import (
	"polling/src/buses/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group(("/buses"))

	addBus := dependencies.AddBusController()
	getAllBus := dependencies.GetAllBusesController()
	updateBus := dependencies.UpdateBusController()
	getBusByIdChofer := dependencies.GetBusByIdChoferController()
	deleteBus := dependencies.DeleteBusController()
	eventBus := dependencies.PollingBusController()

	routes.POST("/", addBus.Run)
	routes.GET("", getAllBus.Run)
	routes.PUT("/:idBus", updateBus.Run)
	routes.GET("/:choferID", getBusByIdChofer.Run)
	routes.DELETE("/:idBus", deleteBus.Run)
	routes.GET("/short", eventBus.ShortPolling)
	routes.GET("/long", eventBus.LongPolling)
}