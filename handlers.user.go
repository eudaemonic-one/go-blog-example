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

func showLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Login"}, "login.html")
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if isUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		render(c, gin.H{
			"title": "Login Success"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle": "Login Failed",
			"ErrorMessage": "Invalid username or password"})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}