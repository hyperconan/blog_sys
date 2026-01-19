package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"user/internal/config"
	"user/internal/models"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Database.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Auto migrate database
	if err := models.MigrateUser(db); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
