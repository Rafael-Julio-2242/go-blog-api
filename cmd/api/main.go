package main

import (
	handler "go-blog-api/internal/handler/posts"
	"go-blog-api/internal/repository"
	"go-blog-api/internal/service"

	"github.com/gin-gonic/gin"
)

// main initializes the Gin router and starts the server on port 3000.

var Storage = make(map[string]string)

func main() {

	Storage["id"] = "0"

	router := gin.Default()

	postRepo := repository.NewPostRepository(Storage)

	postService := service.NewPostService(*postRepo)

	postsRouter := router.Group("/posts")

	{
		postsRouter.GET("/get/:id", handler.NewGetPostHandler(postService).Execute)
		postsRouter.POST("/post", handler.NewCreatePostHandler(postService).Execute)
		postsRouter.PUT("/update", handler.NewUpdatePostHandler(postService).Execute)
		postsRouter.DELETE("/delete/:id", handler.NewDeletePostHandler(postService).Execute)
	}

	router.Run(":3000")
}
