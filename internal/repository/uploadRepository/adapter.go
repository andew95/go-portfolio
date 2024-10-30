package uploadRepository

import (
	"time"

	"gorm.io/gorm"
)

type UploadRepository interface {
	Create(request Upload) (*Upload, error)
	Find() ([]Upload, error)
}

type uploadRepository struct {
	DB *gorm.DB
}

func NewUploadRepository(db *gorm.DB) UploadRepository {
	return &uploadRepository{
		DB: db,
	}
}

func (repo *uploadRepository) Create(request Upload) (*Upload, error) {
	record := &Upload{
		Id:               request.Id,
		FileName:         request.FileName,
		FileType:         request.FileType,
		FilePath:         request.FilePath,
		OriginalFileName: request.OriginalFileName,
		CreatedAt:        time.Now(),
	}
	result := repo.DB.Debug().Create(&record)

	if result.Error != nil {
		return nil, result.Error
	}

	return record, nil
}

func (repo *uploadRepository) Find() ([]Upload, error) {
	var records []Upload
	result := repo.DB.Debug().Find(&records, nil).Order("id DESC")

	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}
