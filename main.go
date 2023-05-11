package main

import (
	"RestLecture/controllers"
	"RestLecture/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.POST("/books/:id/publish", controllers.PublishBook)
	r.DELETE("/books/:id/publish", controllers.UnPublishBook)
	r.POST("auth/register", controllers.RegisterUser)

	err := r.Run()
	if err != nil {
		return
	}
}
