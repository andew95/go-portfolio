package app

import (
	"portfolio/config"
	"portfolio/internal/controller/getHistoryController"
	"portfolio/internal/controller/getSkillController"
	"portfolio/internal/controller/getUploadController"
	"portfolio/internal/controller/uploadController"
)

type AppController struct {
	// AppService       AppService
	UploadController    uploadController.UploadController
	GetUploadController getUploadController.GetUploadController

	GetSkillController getSkillController.GetSkillController

	GetHistoryController getHistoryController.GetHistoryController
}

func SetupHandler(conf config.Config, appService AppService) AppController {
	uploadController := uploadController.NewUploadController(conf, appService.UploadService)
	getUploadController := getUploadController.NewGetUploadController(appService.GetUploadService)

	getSkillController := getSkillController.NewGetHistoryController(appService.GetSkillService)

	getHistoryController := getHistoryController.NewGetHistoryController(appService.GetHistoryService)
	return AppController{
		// AppService:       appService,
		UploadController:     uploadController,
		GetUploadController:  getUploadController,
		GetSkillController:   getSkillController,
		GetHistoryController: getHistoryController,
	}
}
