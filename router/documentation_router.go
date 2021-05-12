package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pdb/database"
	"pdb/model"
)

func CreateDocumentation(c *gin.Context) {
	var documentationData model.DocumentationData

	if err := c.ShouldBindJSON(&documentationData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Documentation"})
		return
	}

	var patient model.Patient
	result := database.GormDB.First(&patient, documentationData.PatientID)
	if result.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	documentation := model.Documentation{DocumentationData: documentationData}
	database.GormDB.Create(&documentation)

	c.JSON(http.StatusOK, documentation)
}

func GetDocumentation(c *gin.Context) {
	id := c.MustGet("id").(int)

	var documentation model.Documentation
	result := database.GormDB.First(&documentation, id)
	if result.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Documentation not found"})
		return
	}

	c.JSON(http.StatusOK, documentation)
}

func ListDocumentation(c *gin.Context) {
	id := c.MustGet("id").(int)

	var documentation []model.Documentation
	database.GormDB.Find(&documentation, "patient_id = ?", id)

	c.JSON(http.StatusOK, documentation)
}
