package config

import (
	"github.com/Gntrstwn19/echo-framework-register-network.git/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=network port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}
