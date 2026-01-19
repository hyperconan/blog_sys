package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"size:200;not null"`
	Content   string    `gorm:"type:text"`
	UserID    uint32    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func MigratePost(db *gorm.DB) error {
	return db.AutoMigrate(&Post{})
}
