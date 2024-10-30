package router

import (
	"portfolio/internal/controller/getUploadController"
	"portfolio/internal/controller/uploadController"

	"github.com/gin-gonic/gin"
)

func SetupUploadRouter(rg *gin.RouterGroup, ctl uploadController.UploadController) {
	rg.POST("", ctl.Execute)
}

func SetupGetUploadRouter(rg *gin.RouterGroup, ctl getUploadController.GetUploadController) {
	rg.GET("", ctl.Execute)
}
