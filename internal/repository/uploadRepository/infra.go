package uploadRepository

import (
	"time"

	"github.com/google/uuid"
)

type Upload struct {
	Id               uuid.UUID `gorm:"primaryKey,"`
	FileName         string
	FileType         string
	FilePath         string
	OriginalFileName string
	CreatedAt        time.Time
}
