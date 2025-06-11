package repository

import (
	model "go-blog-api/internal/model/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (Ur *UserRepository) CreateUser(createUserDTO model.CreateUserDTO) (*model.ResponseCreatedUserDTO, error) {

	user := model.User{
		Name:     createUserDTO.Name,
		Email:    createUserDTO.Email,
		Password: createUserDTO.Password,
	}

	createResult := Ur.db.Create(&user)

	if createResult.Error != nil {
		return nil, createResult.Error
	}

	respUser := model.ResponseCreatedUserDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return &respUser, nil
}

func (Ur *UserRepository) GetUsers() ([]model.ResponseCreatedUserDTO, error) {

	var users []model.User

	result := Ur.db.Model(&model.User{}).Select("*").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	var formatedUsers []model.ResponseCreatedUserDTO

	for _, u := range users {
		formatedUser := model.ResponseCreatedUserDTO{
			Id:        u.Id,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
		}

		formatedUsers = append(formatedUsers, formatedUser)
	}

	return formatedUsers, nil
}
