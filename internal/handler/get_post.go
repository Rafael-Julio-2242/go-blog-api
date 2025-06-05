package handler

import (
	"go-blog-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPostHandler struct {
	postRepo repository.PostRepository
}

func NewGetPostHandler(postRepo repository.PostRepository) *GetPostHandler {
	return &GetPostHandler{
		postRepo: postRepo,
	}
}

func (Gph *GetPostHandler) Execute(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
	}

	post, err := Gph.postRepo.GetPost(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
