package getUploadService

import (
	"portfolio/internal/repository/uploadRepository"
)

type GetUploadService interface {
	Execute() ([]Response, error)
}

type getUploadService struct {
	repo uploadRepository.UploadRepository
}

func NewUploadService(repo uploadRepository.UploadRepository) GetUploadService {
	return &getUploadService{
		repo: repo,
	}
}

func (ser *getUploadService) Execute() ([]Response, error) {
	records, err := ser.repo.Find()
	if err != nil {
		return nil, err
	}

	var result []Response

	for _, record := range records {
		result = append(result, Response{
			Id:               record.Id,
			FileName:         record.FileName,
			FileType:         record.FileType,
			FilePath:         record.FilePath,
			OriginalFileName: record.OriginalFileName,
			CreatedAt:        &record.CreatedAt,
		})
	}
	return result, nil
}
