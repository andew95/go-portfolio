package router

import (
	"portfolio/app"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, ctl app.AppController) {
	r.Static("/static", "./public")

	uploadGroup := r.Group("upload")
	SetupUploadRouter(uploadGroup, ctl.UploadController)
	SetupGetUploadRouter(uploadGroup, ctl.GetUploadController)

	skillGroup := r.Group("skill")
	SetupSkillRouter(skillGroup, ctl.GetSkillController)

	historyGroup := r.Group("history")
	SetupHistoryRouter(historyGroup, ctl.GetHistoryController)
}
