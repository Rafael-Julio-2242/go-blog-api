package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	model "go-blog-api/internal/model/posts"
	"strconv"
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
	storage map[string]string
}

func NewPostRepository(storage map[string]string) *PostRepository {
	return &PostRepository{
		storage: storage,
	}
}

func (Pr *PostRepository) CreatePost(createPostDTO model.CreatePostDTO) (*model.Post, error) {

	nId, _ := strconv.ParseInt(Pr.storage["id"], 10, 64)

	id := fmt.Sprintf("%d", nId+1)

	status := func() string {
		if createPostDTO.Status != nil {
			return *createPostDTO.Status
		}
		return "draft"
	}()

	post := model.Post{
		Id:               id,
		Title:            createPostDTO.Title,
		Summary:          createPostDTO.Summary,
		Content:          createPostDTO.Content,
		Author:           createPostDTO.Author,
		Status:           status,
		Publication_date: "",
	}

	jsonBytes, err := json.Marshal(post)

	if err != nil {
		return nil, err
	}

	Pr.storage[id] = string(jsonBytes)
	Pr.storage["id"] = id

	return &post, nil
}

func (Pr *PostRepository) GetPost(postId string) (*model.Post, error) {

	jsonString := Pr.storage[postId]

	if jsonString == "" {
		return nil, nil
	}

	jsonBytes := []byte(jsonString)

	var post model.Post

	err := json.Unmarshal(jsonBytes, &post)

	if err != nil {
		return nil, errors.New("error on converting from json")
	}

	return &post, nil
}

func (Pr *PostRepository) UpdatePost(updatePostDTO model.UpdatePostDTO) (*model.Post, error) {

	jsonString := Pr.storage[updatePostDTO.Id]

	if jsonString == "" {
		return nil, errors.New("post not found")
	}

	jsonBytes := []byte(jsonString)

	var post model.Post

	err := json.Unmarshal(jsonBytes, &post)

	if err != nil {
		return nil, err
	}

	if updatePostDTO.Title != nil {
		post.Title = *updatePostDTO.Title
	}

	if updatePostDTO.Summary != nil {
		post.Summary = *updatePostDTO.Summary
	}

	if updatePostDTO.Author != nil {
		post.Author = *updatePostDTO.Author
	}

	if updatePostDTO.Content != nil {
		post.Content = *updatePostDTO.Content
	}

	if updatePostDTO.Status != nil {
		post.Status = *updatePostDTO.Status
	}

	if updatePostDTO.Publication_date != nil {
		post.Publication_date = *updatePostDTO.Publication_date
	}

	jsonBytes, err = json.Marshal(post)

	if err != nil {
		return nil, err
	}

	Pr.storage[post.Id] = string(jsonBytes)

	return &post, nil
}

func (Pr *PostRepository) DeletePost(postId string) error {

	Pr.storage[postId] = ""

	return nil
}
