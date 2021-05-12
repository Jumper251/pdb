package router

import (
	"github.com/gin-gonic/gin"
	"pdb/middleware"
)

func Initialize() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	// Patient endpoints
	api.POST("/patients", middleware.ValidatePatient(), CreatePatient)
	api.GET("/patients/:id", middleware.ValidateID(), GetPatient)
	api.DELETE("/patients/:id", middleware.ValidateID(), DeletePatient)
	api.PUT("/patients/:id", middleware.ValidateID(), middleware.ValidatePatient(), UpdatePatient)
	api.GET("/patients", ListPatients)

	// Documentation endpoints
	api.POST("/documentation/:id", middleware.ValidateID(), middleware.PatientExists(), CreateDocumentation)
	api.GET("/documentation/:id", middleware.ValidateID(), GetDocumentation)
	api.GET("/documentation/:id/all", middleware.ValidateID(), middleware.PatientExists(), ListDocumentation)

	return router
}
