package router

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"pdb/middleware"
	"pdb/utils"
	"strings"
)

func Initialize() *gin.Engine {
	router := gin.Default()
	config := utils.GetConfig()

	router.Use(gin.BasicAuth(gin.Accounts{config.AuthUser: config.AuthPassword}))

	// Rewriting for static files
	router.GET("/patients/*p", rewritePath(router))
	router.GET("/documentation/*p", rewritePath(router))
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	router.Use(middleware.Cors())

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

func rewritePath(root *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/api") && c.Request.URL.Path != "/" {
			c.Request.URL.Path = "/"
			root.HandleContext(c)
			c.Abort()
		}
	}
}
