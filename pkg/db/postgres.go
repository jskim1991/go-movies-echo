package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
