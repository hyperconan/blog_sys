package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"size:50;uniqueIndex;not null"`
	Email     string    `gorm:"size:100;not null"`
	Password  string    `gorm:"size:200;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func MigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
