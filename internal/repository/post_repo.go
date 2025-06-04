package repository

import (
	"encoding/json"
	"fmt"
	"go-blog-api/internal/model"
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

func (Pr *PostRepository) CreatePost(createPostDTO model.CreatePostDTO) (*model.Post, error) {

	nId, _ := strconv.ParseInt(Pr.storage["id"], 10, 64)

	id := fmt.Sprintf("%d", nId+1)

	post := model.Post{
		Id:               id,
		Title:            createPostDTO.Title,
		Summary:          createPostDTO.Summary,
		Content:          createPostDTO.Content,
		Author:           createPostDTO.Author,
		Status:           "draft",
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
