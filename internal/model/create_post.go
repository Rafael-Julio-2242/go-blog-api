package model

type CreatePostDTO struct {
	Title   string  `json:"title" binding:"required"`
	Summary string  `json:"summary" binding:"required"`
	Content string  `json:"content" binding:"required"`
	Author  string  `json:"author" binding:"required"`
	Status  *string `json:"status"`
}
