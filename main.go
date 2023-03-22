package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/rest-api-gin/controllers/bookcontroller"
	"github.com/haviz000/rest-api-gin/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/books", bookcontroller.Index)
	r.GET("/api/book/:id", bookcontroller.Show)
	r.POST("/api/book", bookcontroller.Create)
	r.PUT("/api/book/:id", bookcontroller.Update)
	r.DELETE("/api/book", bookcontroller.Delete)

	r.Run()
}
