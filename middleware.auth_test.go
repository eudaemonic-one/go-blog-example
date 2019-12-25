package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), ensureLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})
}

func TestEnsureLoggedInAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), ensureLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestEnsureNotLoggedInUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), ensureNotLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestSetUserStatusAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Fail()
		}
	})
	w := httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	r.ServeHTTP(w, req)
}

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}