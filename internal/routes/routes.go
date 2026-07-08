package routes

import (
	"blog-api/internal/handlers"
	"context"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, ctx context.Context, handle *handlers.HandlePost) {
	mux.HandleFunc("GET /posts", handle.ListHandlerPost(ctx))
	mux.HandleFunc("POST /posts", handle.CreateHandlerPost(ctx))
	mux.HandleFunc("PATCH /posts/{id}", handle.UpdateHandlePost(ctx))
	mux.HandleFunc("GET /posts/{id}", handle.GetHandlerPost(ctx))
	mux.HandleFunc("DELETE /posts/{id}", handle.DeleteHandlePost(ctx))
}
