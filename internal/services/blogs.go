package services

import (
	"blog-api/internal/models"
	"blog-api/internal/repository"
	"context"
	"fmt"
)

type PostRepository interface {
	Create(ctx context.Context, t models.Posts) (models.Posts, error)
	List(ctx context.Context) ([]models.Posts, error)
	Update(ctx context.Context, t models.Posts, id string) (models.Posts, error)
	GetByID(ctx context.Context, t models.Posts, id string) (models.Posts, error)
	Delete(ctx context.Context, t models.Posts, id string) (models.Posts, error)
	Search(ctx context.Context, filter models.PostFilter) ([]models.Posts, error)
}

type PostService struct {
	repo PostRepository
}

func NewPostService(repo *repository.PostRepo) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context, t models.Posts) (models.Posts, error) {
	if t.Title == "" {
		return models.Posts{}, fmt.Errorf("Title is required!")
	}
	return s.repo.Create(ctx, t)
}

func (s *PostService) ListPosts(ctx context.Context) ([]models.Posts, error) {
	return s.repo.List(ctx)
}

func (s *PostService) UpdatePosts(ctx context.Context, t models.Posts, id string) (models.Posts, error) {
	if t.Title == "" && t.Content == "" && t.Category == "" && t.Tags == nil {
		return models.Posts{}, fmt.Errorf("All empty fiels is not Allowed")
	}
	return s.repo.Update(ctx, t, id)
}

func (s *PostService) GetPost(ctx context.Context, t models.Posts, id string) (models.Posts, error) {
	return s.repo.GetByID(ctx, t, id)
}

func (s *PostService) DeletePost(ctx context.Context, t models.Posts, id string) (models.Posts, error) {
	return s.repo.Delete(ctx, t, id)
}

func (s *PostService) Search(ctx context.Context, filter models.PostFilter) ([]models.Posts, error) {
	return s.repo.Search(ctx, filter)
}
