package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("postgres://postgres:postgrespw@localhost:49153"), &gorm.Config{})
}
