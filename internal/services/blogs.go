package services

import (
	"blog-api/internal/models"
	"blog-api/internal/repository"
	"context"
	"fmt"
)

type PostService struct {
	repo *repository.PostRepo
}

func NewPostService(repo *repository.PostRepo) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context, t models.Posts) (models.Posts, error) {
	if t.Title == "" {
		fmt.Println("Title is required!")
	}
	return s.repo.Create(ctx, t)
}

func (s *PostService) ListPosts(ctx context.Context) ([]models.Posts, error) {
	return s.repo.List(ctx)
}
