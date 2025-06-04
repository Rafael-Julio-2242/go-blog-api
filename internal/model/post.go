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

*/

type Post struct {
	Id               string
	Title            string
	Summary          string
	Content          string
	Author           string
	Status           string
	Publication_date string
}
