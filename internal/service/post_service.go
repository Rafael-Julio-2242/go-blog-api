package service

import (
	model "go-blog-api/internal/model/posts"
	"go-blog-api/internal/repository"
)

type PostService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (Ps *PostService) CreatePost(createPostDTO model.CreatePostDTO) (*model.Post, error) {
	return Ps.postRepo.CreatePost(createPostDTO)
}

func (Ps *PostService) GetPosts() ([]model.Post, error) {
	return Ps.postRepo.GetPosts()
}

func (Ps *PostService) UpdatePost(updatePostDTO model.UpdatePostDTO) (*model.Post, error) {
	return Ps.postRepo.UpdatePost(updatePostDTO)
}

func (Ps *PostService) GetPost(id string) (*model.Post, error) {
	return Ps.postRepo.GetPost(id)
}

func (Ps *PostService) DeletePost(id string) error {
	return Ps.postRepo.DeletePost(id)
}
