package model

import (
	"time"

	"github.com/google/uuid"
)

type PocketMessage struct {
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" form:"deletedAt"`
	UUID      uuid.UUID `json:"uuid" gorm:"primaryKey;type=uuid"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	UserUUID  uuid.UUID `json:"userUuid" form:"userUuid"`
}

type Tabler interface {
	TableName() string
}

func (PocketMessage) TableName() string {
	return "pocketMessages"
}
