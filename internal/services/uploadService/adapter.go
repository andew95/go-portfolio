package uploadService

import (
	"portfolio/config"
	"portfolio/internal/repository/uploadRepository"

	"github.com/google/uuid"
)

type UploadServiceAdapter interface {
	Execute(request Request) (*Response, error)
}

type uploadService struct {
	repo uploadRepository.UploadRepository
}

func NewUploadService(conf config.Config, repo uploadRepository.UploadRepository) UploadServiceAdapter {
	return &uploadService{
		repo: repo,
	}
}

func (ser *uploadService) Execute(request Request) (*Response, error) {
	id := uuid.New()

	file, err := request.File.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileName := Rename(request.File.Filename, id.String())
	fileType := request.File.Header.Get("Content-Type")

	err = SaveFileV1(file, fileName)
	if err != nil {
		return nil, err
	}

	record := uploadRepository.Upload{
		Id:               id,
		FileName:         fileName,
		FileType:         fileType,
		FilePath:         "public/" + fileName,
		OriginalFileName: request.File.Filename,
	}

	result, err := ser.repo.Create(record)
	if err != nil {
		return nil, err
	}

	return &Response{
		Id:               result.Id,
		FileName:         result.FileName,
		FileType:         result.FileType,
		FilePath:         result.FilePath,
		OriginalFileName: result.OriginalFileName,
	}, nil
}
