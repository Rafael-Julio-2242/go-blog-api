package repository

import (
	"fmt"
	model "go-blog-api/internal/model/posts"

	"gorm.io/gorm"
)

/*
- Post Structure

id
title
summary
content
author
status "draft" | "posted"
publication_date

*/

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db,
	}
}

func (Pr *PostRepository) CreatePost(createPostDTO model.CreatePostDTO) (*model.Post, error) {

	status := func() string {
		if createPostDTO.Status != nil {
			return *createPostDTO.Status
		}
		return "draft"
	}()

	post := model.Post{
		Title:   createPostDTO.Title,
		Summary: createPostDTO.Summary,
		Content: createPostDTO.Content,
		Author:  createPostDTO.Author,
		Status:  status,
	}

	createResult := Pr.db.Create(&post)

	if createResult.Error != nil {
		return nil, createResult.Error
	}

	return &post, nil
}

func (Pr *PostRepository) GetPosts() ([]model.Post, error) {

	var posts []model.Post

	result := Pr.db.Model(&model.Post{}).Select("*").Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (Pr *PostRepository) GetPost(postId string) (*model.Post, error) {

	var post *model.Post

	result := Pr.db.First(&post, "id = ?", postId)

	if result.RowsAffected <= 0 {
		return nil, nil
	}

	if result.Error != nil {
		fmt.Println("error: ", result.Error)
		return nil, result.Error
	}

	return post, nil
}

func (Pr *PostRepository) UpdatePost(updatePostDTO model.UpdatePostDTO) (*model.Post, error) {

	var post model.Post

	result := Pr.db.First(&post).Where("id = ?", post.Id)

	if result.Error != nil {
		return nil, result.Error
	}

	if updatePostDTO.Content != nil {
		post.Content = *updatePostDTO.Content
	}

	if updatePostDTO.Author != nil {
		post.Author = *updatePostDTO.Author
	}

	if updatePostDTO.Status != nil {
		post.Status = *updatePostDTO.Status
	}

	if updatePostDTO.Summary != nil {
		post.Summary = *updatePostDTO.Summary
	}

	if updatePostDTO.Title != nil {
		post.Title = *updatePostDTO.Title
	}

	if updatePostDTO.Publication_date != nil {
		post.Publication_date = updatePostDTO.Publication_date
	}

	Pr.db.Save(&post)

	return &post, nil
}

func (Pr *PostRepository) DeletePost(postId string) error {

	result := Pr.db.Delete(&model.Post{}, "id = ?", postId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
