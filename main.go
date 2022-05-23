package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default() // Creating a router

	router.LoadHTMLGlob("templates/*") // Loading template files located in 'templates' folder

	initializeRoutes() // Registering routes

	router.Run() // This starts the application on localhost and serves on the 8080 port by default.
}
