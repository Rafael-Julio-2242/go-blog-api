package model

/*
- Post Structure

id
title
summary
content
author
status "draft" | "posted"
publication_date
created_at
user_id

*/

type Post struct {
	Id               int64
	Title            string
	Summary          string
	Content          string
	Author           string
	Status           string
	Publication_date *string
	CreatedAt        string
	UserId           int64
}
