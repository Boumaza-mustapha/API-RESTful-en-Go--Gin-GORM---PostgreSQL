package main

import (
	"GO_CRUD_API/controllers"
	"GO_CRUD_API/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EnvLoadVariables()
	initializers.ConnectTodatabase()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:3000
}
