package main

import "github.com/gin-gonic/gin"

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{
		"title": "Home Page",
		"payload": articles}, "index.html")
}