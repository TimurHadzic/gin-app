package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode) // Sets Gin to Test Mode
	os.Exit(m.Run())          // Runs the other tests
}

func getRouter(withTemplates bool) *gin.Engine { // Helper function to create a router during testing
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) { // Helper function to process a request and test its response

	w := httptest.NewRecorder() // Creates a response recorder
	r.ServeHTTP(w, req)         // Creates the service and process the above request.

	if !f(w) {
		t.Fail()
	}
}

func saveLists() { // This function is used to store the main lists into the temporary one
	tmpArticleList = articleList
}

func restoreLists() { // This function is used to restore the main lists from the temporary one
	articleList = tmpArticleList
}
