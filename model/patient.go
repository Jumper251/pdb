package model

import "gorm.io/gorm"

type PatientData struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type Patient struct {
	gorm.Model
	PatientData
}
