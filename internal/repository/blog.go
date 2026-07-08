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

func (s *PostRepo) Update(ctx context.Context, t models.Posts, id string) (models.Posts, error) {
	result, err := s.db.Exec(ctx, "UPDATE posts SET title = COALESCE(NULLIF($1,''), title), content = COALESCE(NULLIF($2,''), content), category = COALESCE(NULLIF($3,''), category), tags = COALESCE($4, tags) WHERE id = $5", t.Title, t.Content, t.Category, t.Tags, id)
	if err != nil {
		return models.Posts{}, fmt.Errorf("\nError to update Post: %v", err)
	}

	rows := result.RowsAffected()
	if rows == 0 {
		return models.Posts{}, fmt.Errorf("\nNo row updated: %v", err)
	}
	return t, err
}

func (s *PostRepo) GetByID(ctx context.Context, t models.Posts, id string) (models.Posts, error) {
	err := s.db.QueryRow(ctx, "SELECT id, title, content, category, tags FROM posts WHERE id = $1", id).Scan(&t.ID, &t.Title, &t.Content, &t.Category, &t.Tags)
	if err != nil {
		return models.Posts{}, fmt.Errorf("\nError to execute query:%v ", err)
	}
	if t.ID == id {
		return t, nil
	}
	return models.Posts{}, fmt.Errorf("\nInvalid id")
}

func (s *PostRepo) Delete(ctx context.Context, t models.Posts, id string) (models.Posts, error) {
	result, err := s.db.Exec(ctx, "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return models.Posts{}, fmt.Errorf("\nError to execute query:%v ", err)
	}
	rows := result.RowsAffected()
	if rows == 0 {
		return models.Posts{}, fmt.Errorf("\nNo row deleted: %v", err)
	}
	return t, err
}

func (s *PostRepo) Search(ctx context.Context, filter models.PostFilter) ([]models.Posts, error) {
	var args []any
	var argPos int = 1
	var clauses []string
	if filter.Query != "" {
		condiction := fmt.Sprintf(" (title ILIKE '%%' || $%d || '%%') OR (content ILIKE '%%' || $%d || '%%') ", argPos, argPos)
		clauses = append(clauses, condiction)
		args = append(args, filter.Query)
		argPos++
	}
	if filter.Category != "" {
		condiction := fmt.Sprintf(" category ILIKE $%d || '%%'", argPos)
		clauses = append(clauses, condiction)
		args = append(args, filter.Category)
		argPos++
	}
	if filter.Tag != "" {
		condiction := fmt.Sprintf("EXISTS (SELECT 1 FROM unnest(tags) t WHERE t ILIKE $%d || '%%')", argPos)
		clauses = append(clauses, condiction)
		args = append(args, filter.Tag)
		argPos++
	}

	//------------------------------------------------------------------
	query := "SELECT id, title, content, category, tags FROM posts "
	if clauses != nil {
		query += " WHERE 1=1 "
		for i := 0; i < len(clauses); i++ {
			query += " AND " + clauses[i]
		}
	}
	//caso n tenha parametros
	rows, err := s.db.Query(ctx, query, args...)
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
