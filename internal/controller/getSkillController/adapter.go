package getSkillController

import (
	"net/http"
	"portfolio/internal/services/getSkillService"

	"github.com/gin-gonic/gin"
)

type GetSkillController interface {
	Execute(c *gin.Context)
}

type getSkillController struct {
	service getSkillService.GetSkillService
}

func NewGetHistoryController(service getSkillService.GetSkillService) GetSkillController {
	return &getSkillController{
		service: service,
	}
}

func (ctl *getSkillController) Execute(c *gin.Context) {
	response, err := ctl.service.Execute()
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, response)
}
