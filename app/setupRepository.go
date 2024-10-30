package app

import (
	"portfolio/internal/repository/uploadRepository"

	"gorm.io/gorm"
)

type AppRepository struct {
	UploadRepository uploadRepository.UploadRepository
}

func SetupRepository(postgres *gorm.DB) AppRepository {
	uploadRepo := uploadRepository.NewUploadRepository(postgres)
	return AppRepository{
		UploadRepository: uploadRepo,
	}
}
