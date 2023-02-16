package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt     time.Time       `json:"createdAt" form:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt" form:"updatedAt"`
	DeletedAt     time.Time       `json:"deletedAt" form:"deletedAt"`
	UUID          uuid.UUID       `json:"uuid" gorm:"primaryKey;type=uuid"`
	Username      string          `json:"username" form:"username"`
	Password      string          `json:"password" form:"password"`
	PocketMessage []PocketMessage `gorm:"foreignKey:UserUUID;references:UUID;type:VARCHAR(191);constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (User) TableName() string {
	return "users"
}
