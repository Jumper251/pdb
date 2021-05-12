package model

import (
	"gorm.io/gorm"
)

type DocumentationData struct {
	//Time time.Time
	Content   string  `json:"content" binding:"required"`
	PatientID int     `json:"patient_id" binding:"required"`
	Patient   Patient `binding:"-" json:"-"`
}

type Documentation struct {
	gorm.Model
	DocumentationData
}
