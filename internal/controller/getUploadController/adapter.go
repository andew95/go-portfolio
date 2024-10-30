package getUploadController

import (
	"net/http"
	"portfolio/internal/services/getUploadService"

	"github.com/gin-gonic/gin"
)

type GetUploadController interface {
	Execute(c *gin.Context)
}

type getUploadController struct {
	service getUploadService.GetUploadService
}

func NewGetUploadController(service getUploadService.GetUploadService) GetUploadController {
	return &getUploadController{
		service: service,
	}
}

func (ctl *getUploadController) Execute(c *gin.Context) {
	response, err := ctl.service.Execute()
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, response)
}
