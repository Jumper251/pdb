package model

import (
	"gorm.io/gorm"
	"time"
)

type DocumentationData struct {
	Time      time.Time `json:"time" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	PatientID int       `json:"patient_id" binding:"required"`
	Patient   Patient   `binding:"-" json:"-"`
}

type Documentation struct {
	gorm.Model
	DocumentationData
}
