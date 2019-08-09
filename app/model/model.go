package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PM25 struct {
	gorm.Model
	Value float64 `json:value`
}

type PM10 struct {
	gorm.Model
	Value float64 `json:value`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&PM25{}, &PM10{})
	return db
}
