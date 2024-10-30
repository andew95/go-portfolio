package app

import (
	"portfolio/config"
	"portfolio/internal/services/getHistoryService"
	"portfolio/internal/services/getSkillService"
	"portfolio/internal/services/getUploadService"
	"portfolio/internal/services/uploadService"
)

type AppService struct {
	AppRepo          AppRepository
	UploadService    uploadService.UploadServiceAdapter
	GetUploadService getUploadService.GetUploadService

	GetSkillService getSkillService.GetSkillService

	GetHistoryService getHistoryService.GetHistoryService
}

func SetupAppSevice(conf config.Config, appRepo AppRepository) AppService {

	uploadService := uploadService.NewUploadService(conf, appRepo.UploadRepository)
	getUploadService := getUploadService.NewUploadService(appRepo.UploadRepository)

	getSkillService := getSkillService.NewGetSkillService()
	getHistoryService := getHistoryService.NewGetHistory()
	return AppService{
		UploadService:    uploadService,
		GetUploadService: getUploadService,

		GetSkillService: getSkillService,

		GetHistoryService: getHistoryService,
	}
}
