package repository

import (
	"blog-api/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(pool *pgxpool.Pool) *PostRepo {
	return &PostRepo{db: pool}
}

func (s *PostRepo) Create(ctx context.Context, t models.Posts) (models.Posts, error) {
	err := s.db.QueryRow(ctx, "INSERT INTO posts(title, content, category, tags) VALUES($1, $2, $3, $4) RETURNING id", &t.Title, &t.Content, &t.Category, &t.Tags).Scan(&t.ID)

	return t, err
}

func (s *PostRepo) List(ctx context.Context) ([]models.Posts, error) {
	rows, err := s.db.Query(ctx, "SELECT id, title, content, category, tags FROM posts")
	if err != nil {
		return nil, fmt.Errorf("\nError to execute query:%v ", err)
	}
	defer rows.Close()

	var post []models.Posts
	for rows.Next() {
		var p models.Posts
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.Tags)
		if err != nil {
			return nil, fmt.Errorf("\nError to find Posts: %v", err)
		}
		post = append(post, p)
	}
	return post, rows.Err()
}
