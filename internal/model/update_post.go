package model

type UpdatePostDTO struct {
	Id               string
	Title            *string
	Summary          *string
	Content          *string
	Author           *string
	Status           *string
	Publication_date *string
}
