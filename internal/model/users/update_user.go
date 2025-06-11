package model

type UpdateUserDTO struct {
	Id       int64
	Email    *string
	Name     *string
	Password *string
}
