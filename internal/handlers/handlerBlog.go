package handlers

import (
	"blog-api/internal/models"
	"blog-api/internal/services"
	"blog-api/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlePost struct {
	svc *services.PostService
}

func NewHandlePost(svc *services.PostService) *HandlePost {
	return &HandlePost{svc: svc}
}

func (s HandlePost) CreateHandlerPost(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Posts

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid data")
			fmt.Printf("\nError to get client data: %v", err)
			return
		}
		created, err := s.svc.CreatePost(ctx, p)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Error to create Post-Something got wrong-Title field is required")
			fmt.Printf("\nError to create Post: %v", err)
			return
		}
		utils.JSONResponse(w, http.StatusOK, created)
	}
}

func (s HandlePost) ListHandlerPost(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		if len(params) == 0 {
			listed, err := s.svc.ListPosts(ctx)
			if err != nil {
				utils.ErrorResponse(w, http.StatusInternalServerError, "Error to find Posts")
				fmt.Printf("\nError to find Post: %v ", err)
				return
			}
			utils.JSONResponse(w, http.StatusOK, listed)
			return
		}

		filter := models.PostFilter{
			Query:    params.Get("q"),
			Category: params.Get("category"),
			Tag:      params.Get("tag"),
		}
		listed, err := s.svc.Search(ctx, filter)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error to find Posts")
			fmt.Printf("\nError to find Post: %v ", err)
			return
		}
		utils.JSONResponse(w, http.StatusOK, listed)
	}
}

func (s HandlePost) UpdateHandlePost(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t models.Posts

		id := r.PathValue("id")

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid data")
			fmt.Printf("\nError to update post: %v", err)
			return
		}
		updated, err := s.svc.UpdatePosts(ctx, t, id)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid data to update your post")
			fmt.Printf("Invalid data to update your post: %v", err)
			return
		}
		utils.JSONResponse(w, http.StatusOK, updated)
	}
}

func (s HandlePost) GetHandlerPost(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t models.Posts
		id := r.PathValue("id")

		findedID, err := s.svc.GetPost(ctx, t, id)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error to find Post")
			fmt.Printf("\nError to find Post: %v ", err)
			return
		}
		utils.JSONResponse(w, http.StatusOK, findedID)
	}
}

func (s HandlePost) DeleteHandlePost(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t models.Posts
		id := r.PathValue("id")

		deleted, err := s.svc.DeletePost(ctx, t, id)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error to delete Post")
			fmt.Printf("\nError to delete Post: %v ", err)
			return
		}
		utils.JSONResponse(w, http.StatusOK, deleted)
	}
}
