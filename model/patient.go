package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type PatientData struct {
	FirstName          string `json:"first_name" binding:"required"`
	LastName           string `json:"last_name" binding:"required"`
	Birthdate          string `json:"birthdate"`
	Street             string `json:"street"`
	ZipCode            string `json:"zip_code"`
	City               string `json:"city"`
	Country            string `json:"country"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	Diagnosis          string `json:"diagnosis"`
	SuspectedDiagnosis string `json:"suspected_diagnosis"`
	Medication         string `json:"medication"`
	Sex                string `json:"sex"`
}

type Patient struct {
	gorm.Model
	PatientData
}

func FindWithQuery(query string, db *gorm.DB) []Patient {
	var patients []Patient
	if len(query) <= 0 {
		db.Find(&patients)
		return patients
	}

	split := strings.Split(query, " ")

	switch len(split) {
	case 1:
		db.Where("first_name LIKE ?", fmt.Sprintf("%%%s%%", split[0])).
			Or("last_name LIKE ?", fmt.Sprintf("%%%s%%", split[0])).
			Find(&patients)
	case 2:
		db.Where("first_name LIKE ?", fmt.Sprintf("%%%s%%", split[0])).
			Where("last_name LIKE ?", fmt.Sprintf("%%%s%%", split[1])).
			Find(&patients)
	default:
		lastSplit := len(split) - 1
		db.Where("first_name LIKE ?", fmt.Sprintf("%%%s%%", strings.Join(split[0:lastSplit], " "))).
			Where("last_name LIKE ?", fmt.Sprintf("%%%s%%", split[lastSplit])).
			Find(&patients)
	}

	return patients

}
