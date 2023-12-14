package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectionDatabase() {
	database, err := gorm.Open(mysql.Open("notes:tmp_pwd@tcp(127.0.0.1:3306)/notes?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic("failed to connected to database!")
	}

	DB = database
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/**/**")

	connectionDatabase()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "NOTE APP",
		})
	})

	log.Println("Server Started!")
	r.Run() // Default Port 8080
}
