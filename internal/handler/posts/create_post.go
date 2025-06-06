package handler

import (
	model "go-blog-api/internal/model/posts"
	"go-blog-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostHandler struct {
	postService *service.PostService
}

func NewCreatePostHandler(postService *service.PostService) *CreatePostHandler {

	return &CreatePostHandler{
		postService: postService,
	}
}
func (Cph *CreatePostHandler) Execute(c *gin.Context) {
	var dto model.CreatePostDTO

	if err := c.ShouldBindBodyWithJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := Cph.postService.CreatePost(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})

}
