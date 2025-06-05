package handler

import (
	"go-blog-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePostHandler struct {
	postRepo repository.PostRepository
}

func NewDeletePostHandler(postRepo repository.PostRepository) *DeletePostHandler {
	return &DeletePostHandler{
		postRepo: postRepo,
	}
}

func (Dph *DeletePostHandler) Execute(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err := Dph.postRepo.DeletePost(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
