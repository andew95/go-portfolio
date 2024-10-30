package getHistoryController

import (
	"net/http"
	"portfolio/internal/services/getHistoryService"

	"github.com/gin-gonic/gin"
)

type GetHistoryController interface {
	Execute(c *gin.Context)
}

type getHistoryController struct {
	service getHistoryService.GetHistoryService
}

func NewGetHistoryController(service getHistoryService.GetHistoryService) GetHistoryController {
	return &getHistoryController{
		service: service,
	}
}

func (ctl *getHistoryController) Execute(c *gin.Context) {
	response, err := ctl.service.Execute()
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, response)
}
