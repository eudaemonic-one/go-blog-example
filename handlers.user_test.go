package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestShowRegisterPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/user/register", showRegisterPage)
	req, _ := http.NewRequest("GET", "/user/register", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0
		return statusOK && pageOK
	})
}

func TestRegisterUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/user/register", register)
	payload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/user/register", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register Success</title>") > 0
		return statusOK && pageOK
	})
}

func TestRegisterUnavailableUsername(t *testing.T) {
	r := getRouter(true)
	r.POST("/user/register", register)
	payload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/user/register", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload)))
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusBadRequest
		return statusOK
	})
}

func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")
	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")
	return params.Encode()
}