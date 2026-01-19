package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint32    `gorm:"not null"`
	PostID    uint32    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func MigrateComment(db *gorm.DB) error {
	return db.AutoMigrate(&Comment{})
}
