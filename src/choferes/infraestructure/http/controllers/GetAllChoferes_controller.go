package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"polling/src/choferes/application"
)

type GetAllChoferesController struct {
	uc *application.GetAllChoferesUseCase
}

func NewGetAllChoferesController(uc *application.GetAllChoferesUseCase) *GetAllChoferesController {
	return &GetAllChoferesController{uc: uc}
}

func (ctrl *GetAllChoferesController) Run(c *gin.Context) {
	choferes, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"choferes": choferes})
}