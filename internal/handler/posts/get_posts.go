package handler

import (
	"fmt"
	"go-blog-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPostsHandler struct {
	postService *service.PostService
}

func NewGetPostsHandler(postService *service.PostService) *GetPostsHandler {
	return &GetPostsHandler{
		postService: postService,
	}
}

func (Gph *GetPostsHandler) Execute(c *gin.Context) {
	posts, err := Gph.postService.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
