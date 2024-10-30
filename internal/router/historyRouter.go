package router

import (
	"portfolio/internal/controller/getHistoryController"

	"github.com/gin-gonic/gin"
)

func SetupHistoryRouter(rg *gin.RouterGroup, ctl getHistoryController.GetHistoryController) {
	rg.GET("", ctl.Execute)
}
