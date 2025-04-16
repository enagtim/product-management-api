package db

import (
	"product-management-api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *DB {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{db}
}
