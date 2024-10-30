package uploadService

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type Request struct {
	File *multipart.FileHeader
}

type Response struct {
	Id               uuid.UUID `json:"id"`
	FileName         string    `json:"fileName"`
	FileType         string    `json:"fileType"`
	FilePath         string    `json:"filePath"`
	OriginalFileName string    `json:"originalFileName"`
}
