package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)


var tmpUserList []user

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		r.Use(setUserStatus())
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {
	req, _ := http.NewRequest("GET", "/", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}

func saveUserList() {
	tmpUserList = userList
}

func restoreUserList() {
	userList = tmpUserList
}