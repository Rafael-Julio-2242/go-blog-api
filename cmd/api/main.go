package main

import (
	"go-blog-api/internal/config"
	handler "go-blog-api/internal/handler/posts"
	"go-blog-api/internal/repository"
	"go-blog-api/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main initializes the Gin router and starts the server on port 3000.

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error on loading .env: ", err)
	}

	db := config.GetConnection(os.Getenv("DATABASE_URL"))

	router := gin.Default()

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(*postRepo)
	postsRouter := router.Group("/posts")

	{
		postsRouter.GET("/get/:id", handler.NewGetPostHandler(postService).Execute)
		postsRouter.GET("/get", handler.NewGetPostsHandler(postService).Execute)
		postsRouter.POST("/post", handler.NewCreatePostHandler(postService).Execute)
		postsRouter.PUT("/update", handler.NewUpdatePostHandler(postService).Execute)
		postsRouter.DELETE("/delete/:id", handler.NewDeletePostHandler(postService).Execute)
	}

	router.Run(":3000")
}
