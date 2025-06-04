package repository

import "go-blog-api/internal/model"

/*
- Post Structure

id
title
resume
content
author
status "draft" | "posted"
publication_date

*/

type PostRepository struct{}

func (Pr *PostRepository) CreatePost(createPostDTO model.CreatePostDTO) {

}
