package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syunsuke-I/gin_notes/models"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectionDatabase()
	models.DbMigrate()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "NOTE APP",
		})
	})

	log.Println("Server Started!")
	r.Run() // Default Port 8080
}
