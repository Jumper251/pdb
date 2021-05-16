package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pdb/database"
	"pdb/model"
)

func CreatePatient(c *gin.Context) {
	patientData := c.MustGet("patientData").(model.PatientData)

	patient := model.Patient{PatientData: patientData}
	database.GormDB.Create(&patient)

	c.JSON(http.StatusCreated, patient)
}

func GetPatient(c *gin.Context) {
	id := c.MustGet("id").(int)

	var patient model.Patient
	result := database.GormDB.First(&patient, id)
	if result.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func DeletePatient(c *gin.Context) {
	id := c.MustGet("id").(int)

	result := database.GormDB.Unscoped().Delete(&model.Patient{}, id)
	if result.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	database.GormDB.Unscoped().Where("patient_id = ?", id).Delete(&model.Documentation{})

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}

func UpdatePatient(c *gin.Context) {
	id := c.MustGet("id").(int)
	patientData := c.MustGet("patientData").(model.PatientData)

	var patient model.Patient
	result := database.GormDB.First(&patient, id)
	if result.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	patient.PatientData = patientData
	database.GormDB.Save(&patient)

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated"})
}

func ListPatients(c *gin.Context) {
	query := c.Query("query")
	patients := model.FindWithQuery(query, database.GormDB)

	c.JSON(http.StatusOK, patients)
}
