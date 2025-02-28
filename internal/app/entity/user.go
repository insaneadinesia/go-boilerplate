package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID          uuid.UUID `gorm:"primaryKey"`
	Name          string
	Username      string
	Email         string
	SubDistrictID int64
	CreatedAt     time.Time `gorm:"<-:create"`
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()
	return
}
