package usecase

import (
	"go_echo_api/model"
	"go_echo_api/repository"
)

type IPostUsecase interface {
	GetAllPosts() ([]model.Post, error)
	GetPostById(postId uint) (model.Post, error)
	CreatePost(post model.Post) (model.Post, error)
	UpdatePost(post model.Post, postId uint) (model.Post, error)
	DeletePost(postId uint) error
}

type postUsecase struct {
	pr repository.IPostRepository
}

func NewPostUsecase(pr repository.IPostRepository) IPostUsecase {
	return &postUsecase{pr}
}

func (pu *postUsecase) GetAllPosts() ([]model.Post, error) {
	posts := []model.Post{}
	if err := pu.pr.GetAllPosts(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (pu *postUsecase) GetPostById(postId uint) (model.Post, error) {
	post := model.Post{}
	if err := pu.pr.GetPostById(&post, postId); err != nil {
		return model.Post{}, err
	}
	resPost := model.Post{
		ID:        post.ID,
		Title:     post.Title,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) CreatePost(post model.Post) (model.Post, error) {
	if err := pu.pr.CreatePost(&post); err != nil {
		return model.Post{}, err
	}
	resPost := model.Post{
		ID:        post.ID,
		Title:     post.Title,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) UpdatePost(post model.Post, postId uint) (model.Post, error) {
	if err := pu.pr.UpdatePost(&post, postId); err != nil {
		return model.Post{}, err
	}
	resPost := model.Post{
		ID:        post.ID,
		Title:     post.Title,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) DeletePost(postId uint) error {
	if err := pu.pr.DeletePost(postId); err != nil {
		return err
	}
	return nil
}
