package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PM25 struct {
	gorm.Model
	Value   float32   `json:value`
	DateUTC time.Time `json:DateUTC`
}

type PM10 struct {
	gorm.Model
	Value   float32   `json:value`
	DateUTC time.Time `json:DateUTC`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&PM25{}, &PM10{})
	return db
}
