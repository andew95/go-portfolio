package uploadRepository

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Id               uuid.UUID `gorm:"primaryKey,"`
	FileName         string
	FileType         string
	FilePath         string
	OriginalFileName string
}

type Response struct {
	Id               uuid.UUID `gorm:"primaryKey,"`
	FileName         string
	FileType         string
	FilePath         string
	OriginalFileName string
	CreatedAt        time.Time
}
