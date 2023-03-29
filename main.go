package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/rest-api-gin/controllers/bookcontroller"
	"github.com/haviz000/rest-api-gin/database"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	r.GET("/api/book/:id", bookcontroller.Show)
	r.GET("/api/books", bookcontroller.Index)
	r.POST("/api/book", bookcontroller.Create)
	r.PUT("/api/book/:id", bookcontroller.Update)
	r.DELETE("/api/book/:id", bookcontroller.Delete)

	r.Run()
}
