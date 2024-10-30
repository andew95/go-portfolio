package getUploadService

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	Id               uuid.UUID  `json:"id"`
	FileName         string     `json:"fileName"`
	FileType         string     `json:"fileType"`
	FilePath         string     `json:"filePath"`
	OriginalFileName string     `json:"originalFileName"`
	CreatedAt        *time.Time `json:"createdAt"`
}
