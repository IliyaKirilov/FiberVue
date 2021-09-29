package models

import (
	"main/go/pkg/mod/go.mongodb.org/mongo-driver@v1.5.3/x/mongo/driver/uuid"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

type Base struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.UUID = uuid.New()

	t := GeGenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil

}

func (base *Base) AfterUpdate(tx *gorm.DB) error {
	base.UpdatedAt = GenerateISOString()

	return nil
}
