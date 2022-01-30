package model

import (
	"github.com/ashhhleyyy/clickd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	dsn := config.Configuration.DbDsn
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	Database = db
}

func InitDB() {
	Database.AutoMigrate(&Event{})
}
