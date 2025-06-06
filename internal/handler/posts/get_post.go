package handler

import (
	"go-blog-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPostHandler struct {
	postService *service.PostService
}

func NewGetPostHandler(postService *service.PostService) *GetPostHandler {
	return &GetPostHandler{
		postService: postService,
	}
}

func (Gph *GetPostHandler) Execute(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
	}

	post, err := Gph.postService.GetPost(id)

	if post == nil && err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post Not Found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
