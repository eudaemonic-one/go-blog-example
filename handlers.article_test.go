package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)
	req, _ := http.NewRequest("GET", "/", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool{
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0
		return statusOK && pageOK
	})
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)
	r.GET("/", showIndexPage)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept","application/json")
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []article
		err = json.Unmarshal(p, &articles)
		return err == nil && statusOK && len(articles) >= 2
	})
}

func TestArticleXML(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var a article
		err = xml.Unmarshal(p, &a)
		return err == nil && statusOK && len(a.Title) >= 0
	})
}