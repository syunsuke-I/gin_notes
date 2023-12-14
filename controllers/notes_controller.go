package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syunsuke-I/gin_notes/models"
)

func NoteIndex(c *gin.Context) {
	notes := models.NotesAll()
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"notes/new.html",
		gin.H{},
	)
}

func NotesCreate(c *gin.Context) {
	name := c.PostForm("name")
	content := c.PostForm("content")

	models.NoteCreate(name, content)

	c.Redirect(
		http.StatusMovedPermanently,
		"notes",
	)
}
