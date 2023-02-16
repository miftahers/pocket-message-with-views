package model

import (
	"time"

	"github.com/google/uuid"
)

type UniqueId struct {
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" form:"deletedAt"`
	RandomID  string    `json:"randomId" form:"randomId"`
	Visit     int       `json:"visit"`
	UUID      uuid.UUID `json:"uuid" form:"uuid" gorm:"type:VARCHAR(191)"`
}

func (UniqueId) TableName() string {
	return "uniqueIds"
}
