package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"rest-apis/gin-rest-api/controllers"
	"rest-apis/gin-rest-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	if len(os.Args) > 1 {
		models.BulkBookLoad(os.Args[1])
	}

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run(":5000")
}
