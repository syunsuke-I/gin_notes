package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syunsuke-I/gin_notes/controllers"
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

	r.GET("/notes", controllers.NoteIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)
	r.DELETE("/notes/:id", controllers.NotesDelete)
	r.DELETE("/notes/", controllers.NoteIndex)
	r.GET("/notes/edit/:id", controllers.NoteGet)
	r.POST("/notes/:id", controllers.NoteEdit)
	log.Println("Server Started!")
	r.Run() // Default Port 8080
}
