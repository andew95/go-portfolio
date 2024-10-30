package router

import (
	"portfolio/internal/controller/getSkillController"

	"github.com/gin-gonic/gin"
)

func SetupSkillRouter(rg *gin.RouterGroup, ctl getSkillController.GetSkillController) {
	rg.GET("", ctl.Execute)
}
