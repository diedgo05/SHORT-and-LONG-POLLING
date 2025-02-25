package controllers

import (
	"net/http"
	"polling/src/buses/application"
	"time"

	"github.com/gin-gonic/gin"
)

type EventBusController struct {
	uc *application.EventBusUseCase
}

func NewEventBusController(uc *application.EventBusUseCase) *EventBusController {
	return &EventBusController{uc: uc}
}

func (ctrl *EventBusController) ShortPolling(c *gin.Context) {
	bus := ctrl.uc.Run()
	if bus == nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "No hay nuevos eventos"})
	} else {
		c.JSON(http.StatusOK, bus)
	}
}

func (ctrl *EventBusController) LongPolling(c *gin.Context) {
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusGatewayTimeout, gin.H{"message": "Timeout: No hay nuevos eventos"})
			return
		case <-ticker.C:
			buses := ctrl.uc.Run()
			if buses != nil {
				c.JSON(http.StatusOK, buses)
				return
			}
		}
	}
}