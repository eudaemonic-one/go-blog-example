package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func showRegisterPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if _, err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c, gin.H{
			"title": "Register Success"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle": "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}