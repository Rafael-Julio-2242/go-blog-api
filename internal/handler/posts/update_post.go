package handler

import (
	model "go-blog-api/internal/model/posts"
	"go-blog-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePostHandler struct {
	postService *service.PostService
}

func NewUpdatePostHandler(postService *service.PostService) *UpdatePostHandler {
	return &UpdatePostHandler{
		postService: postService,
	}
}

func (Uph *UpdatePostHandler) Execute(c *gin.Context) {
	var dto model.UpdatePostDTO

	if err := c.ShouldBindBodyWithJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dto.Author == nil &&
		dto.Content == nil &&
		dto.Publication_date == nil &&
		dto.Summary == nil &&
		dto.Title == nil &&
		dto.Status == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nothing to update"})
		return
	}

	post, err := Uph.postService.UpdatePost(dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
