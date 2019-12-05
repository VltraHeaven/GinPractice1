package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Instantiate gin router
	router := gin.Default()

	// Create a slice of filepaths
	rt := []string{
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/menu.html",
	}

	// Instantiate elements from filepath slice and pass to router
	idx := rt[0]
	hdr := rt[1]
	ftr := rt[2]
	nv := rt[3]

	router.LoadHTMLFiles(idx, hdr, ftr, nv)

	// Set routes
	router.GET("/", func(c *gin.Context) {

		//Call HTML method of Context to render a template
		//Set HTTP status to 200(OK)
		//Use 'index.html' template
		//Pass data from that the page uses
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	})

	// Serve the web application
	router.Run(":9090")
}
