package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rynfkn/crud-app/controllers"
	"github.com/rynfkn/crud-app/initializers"
	"github.com/rynfkn/crud-app/migrate"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// migrate
	migrate.Migrate()
	
	r := gin.Default()
	
	r.GET("/post", controllers.GetAllPost)
	r.GET("/post/:id", controllers.GetPostById)

	r.POST("/post", controllers.AddNewPost)

	r.PUT("/post/:id", controllers.UpdatePostById)

	r.DELETE("/post/:id", controllers.DeletePostById)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}