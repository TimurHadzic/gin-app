package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(c, gin.H{ // Calls the render function with the name of the template to render
		"title":   "Home Page",
		"payload": articles}, "index.html")

}

func render(c *gin.Context, data gin.H, templateName string) { // Renders one of HTML, JSON or CSV based on the 'Accept' header of the request, iff the header doesn't specify this, HTML is rendered

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"]) // Responds with JSON
	case "application/xml":
		c.XML(http.StatusOK, data["payload"]) // Responds with XML
	default:
		c.HTML(http.StatusOK, templateName, data) // Responds with HTML
	}

}

func getArticle(c *gin.Context) {

	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil { // Checks if the article ID is valid
		if article, err := getArticleByID(articleID); err == nil { // Checks if the article exists
			c.HTML( // Calls the HTML method of the Context to render a template
				http.StatusOK,  // Sets the HTTP status to 200 (OK)
				"article.html", // Uses the index.html template
				gin.H{ // Pass the data that the page uses
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err) // If the article is not found, aborts with an error
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound) // If an invalid article ID is specified in the URL, aborts with an error
	}
}
