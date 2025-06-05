package handler

import (
	"go-blog-api/internal/model"
	"go-blog-api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostHandler struct {
	postRepo repository.PostRepository
}

func NewCreatePostHandler(postRepo repository.PostRepository) *CreatePostHandler {

	return &CreatePostHandler{
		postRepo: postRepo,
	}
}
func (Cph *CreatePostHandler) Execute(c *gin.Context) {
	var dto model.CreatePostDTO

	if err := c.ShouldBindBodyWithJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := Cph.postRepo.CreatePost(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})

}
