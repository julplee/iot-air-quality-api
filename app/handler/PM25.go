package handler

import (
	"net/http"

	"../model"
	"github.com/jinzhu/gorm"
)

func GetAllPM25(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	PM25s := []model.PM25{}
	db.Find(&PM25s)
	respondJSON(w, http.StatusOK, PM25s)
}
