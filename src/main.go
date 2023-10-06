package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	log.Println("booting server...")
	
	router := gin.Default()

	router.LoadHTMLGlob("src/templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/bruh", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message": "hello" })
	})

	router.GET("/rawhtml", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte("<html><head><title>FROM GOLANG</title></head><body>yah brah</body></html>"))
	})

	router.Run()
}