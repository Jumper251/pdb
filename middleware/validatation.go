package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pdb/database"
	"pdb/model"
	"strconv"
)

func ValidateID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		c.Set("id", id)

		c.Next()
	}
}

func ValidatePatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patientData model.PatientData

		if err := c.ShouldBindJSON(&patientData); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Patient"})
			return
		}

		c.Set("patientData", patientData)

		c.Next()
	}
}

func PatientExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("id").(int)

		var patient model.Patient
		result := database.GormDB.First(&patient, id)
		if result.RowsAffected <= 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		c.Set("patient", patient)

		c.Next()
	}
}
