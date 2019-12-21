package main

import "github.com/gin-gonic/gin"

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