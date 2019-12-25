package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)
	// Set the router as the default one provided by Gin
	router = gin.Default()
	// Load templates before the start
	router.LoadHTMLGlob("templates/*")
	// Initialize the routes
	initializeRoutes()
	// Start serving the application
	router.Run()
}

func render(c *gin.Context, data gin.H, tenplateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, tenplateName, data)
	}
}