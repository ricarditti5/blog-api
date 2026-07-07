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
		created, err := s.svc.ListPosts(ctx)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Error to find Posts")
			fmt.Printf("\nError to find Post: %v ", err)
			return
		}

		utils.JSONResponse(w, http.StatusOK, created)
	}
}
