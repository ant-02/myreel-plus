package model

import (
	"time"

	"github.com/ant-02/myreel-plus/pkg/constants"
	"gorm.io/gorm"
)

type UserCredential struct {
	ID             int64          `gorm:"column:id;type:bigint;primary_key"`
	Email          string         `gorm:"column:email;type:varchar(255);uniqueIndex;not null"`
	Phone          string         `gorm:"column:phone;type:varchar(20);uniqueIndex"`
	PasswordHash   string         `gorm:"column:password_hash;type:varchar(255);not null"`
	Status         int64          `gorm:"column:status;type:int;default:1;index"`
	FailedAttempts int64          `gorm:"column:failed_attempts;type:int;default:0"`
	LockedUntil    *time.Time     `gorm:"column:locked_until;type:timestamp"`
	LastLoginAt    *time.Time     `gorm:"column:last_login_at;type:timestamp"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;index"`
}

func (UserCredential) TableName() string {
	return constants.UserCredentialName
}
