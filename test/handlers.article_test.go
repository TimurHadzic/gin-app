package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testShowIndexPageUnauthenticated(t *testing.T) { // Test that a GET request to the home page returns the home page with the HTTP code 200 for an unauthenticated user
	r := getRouter(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil) // Create a request to send to the above route

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>home page</title>") > 0 // Test that the page title is "Home Page"

		return statusOK && pageOK
	})
}
