package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	// Instantiate gin router
	router := gin.Default()

	router.LoadHTMLFiles("templates/index.html")
	router.GET("/", func(c *gin.Context) {
		//Call HTML method of Context to render a template
		//Set HTTP status to 200(OK)
		//Use 'index.html' template
		//Pass data from that the page uses
		c.HTML(http.StatusOK, "index.html", gin.H{ "title": "Home Page", },
			)
	})
	// Serve the web application
	router.Run(":9090")
}