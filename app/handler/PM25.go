package handler

import (
	"encoding/json"
	"net/http"

	"../model"
	"gorm.io/gorm"
)

func GetAllPM25(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	PM25s := []model.PM25{}
	db.Find(&PM25s)
	respondJSON(w, http.StatusOK, PM25s)
}

func CreatePM25(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	PM25 := model.PM25{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&PM25); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&PM25).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, PM25)
}
