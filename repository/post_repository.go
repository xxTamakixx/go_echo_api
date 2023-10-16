package repository

import (
	"fmt"
	"go_echo_api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPostRepository interface {
	GetAllPosts(posts *[]model.Post) error
	GetPostById(post *model.Post, postId uint) error
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post, postId uint) error
	DeletePost(postId uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{db}
}

func (pr *postRepository) GetAllPosts(posts *[]model.Post) error {
	if err := pr.db.Order("created_at").Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetPostById(post *model.Post, postId uint) error {
	if err := pr.db.First(post, postId).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) CreatePost(post *model.Post) error {
	if err := pr.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) UpdatePost(post *model.Post, postId uint) error {
	result := pr.db.Model(post).Clauses(clause.Returning{}).Where("id=?", postId).Update("title", post.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (pr *postRepository) DeletePost(postId uint) error {
	reslut := pr.db.Where("id=?", postId).Delete(&model.Post{})
	if reslut.Error != nil {
		return reslut.Error
	}
	if reslut.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
