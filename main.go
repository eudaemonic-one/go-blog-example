package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
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
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, tenplateName, data)
	}
}