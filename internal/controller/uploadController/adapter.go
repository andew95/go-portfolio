package uploadController

import (
	"errors"
	"net/http"
	"portfolio/config"
	"portfolio/internal/services/uploadService"

	"github.com/gin-gonic/gin"
)

type UploadController interface {
	Execute(c *gin.Context)
}

type uploadHandlerAdapter struct {
	conf    config.Config
	service uploadService.UploadServiceAdapter
}

func NewUploadController(conf config.Config, uploadService uploadService.UploadServiceAdapter) UploadController {
	return &uploadHandlerAdapter{
		conf:    conf,
		service: uploadService,
	}
}

func (ctl *uploadHandlerAdapter) Execute(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
	}
	if file == nil {
		c.Error(errors.New("file not found"))
	}

	request := uploadService.Request{
		File: file,
	}

	response, err := ctl.service.Execute(request)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, response)
}
