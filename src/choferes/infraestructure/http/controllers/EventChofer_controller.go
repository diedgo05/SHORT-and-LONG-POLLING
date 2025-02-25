package controllers

import (
	"net/http"
	"polling/src/choferes/application"
	"time"

	"github.com/gin-gonic/gin"
)

type EventChoferController struct {
	uc *application.EventChoferUseCase
}

func NewEventChoferController(uc *application.EventChoferUseCase) *EventChoferController {
	return &EventChoferController{uc: uc}
}


func (ctrl *EventChoferController) ShortPolling(c *gin.Context) {
	chofer := ctrl.uc.Run()
	if chofer == nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "No hay nuevos eventos"})
	} else {
		c.JSON(http.StatusOK, chofer)
	}
}


func (ctrl *EventChoferController) LongPolling(c *gin.Context) {
	timeout := time.After(30 * time.Second)  
	ticker := time.NewTicker(1 * time.Second) 
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusGatewayTimeout, gin.H{"message": "Timeout: No hay nuevos eventos"})
			return
		case <-ticker.C:
			choferes := ctrl.uc.Run() 
			if choferes != nil {
				c.JSON(http.StatusOK, choferes)
				return
			}
		}
	}
}
