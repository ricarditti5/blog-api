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
	var p models.Posts
	err := s.db.QueryRow(ctx, "INSERT INTO posts(title, content, category, tags) VALUES($1, $2, $3, $4) RETURNING id", p.Title, p.Content, p.Category, p.Tags).Scan(&p.ID)

	return p, err
}

func (s *PostRepo) List(ctx context.Context) ([]models.Posts, error) {
	rows, err := s.db.Query(ctx, "SELECT title, content, category, tags FROM posts")
	if err != nil {
		return nil, fmt.Errorf("Error to execute query:%v ", err)
	}
	defer rows.Close()

	var post []models.Posts
	for rows.Next() {
		var p models.Posts
		err := rows.Scan(&p.Title, &p.Content, &p.Category, &p.Tags)
		if err != nil {
			fmt.Println("Error to find Posts")
			return nil, err
		}
		post = append(post, p)
	}
	return post, rows.Err()
}
