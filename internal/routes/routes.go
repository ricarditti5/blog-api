package routes

import (
	"blog-api/internal/handlers"
	"context"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, ctx context.Context, handle *handlers.HandlePost) {
	mux.HandleFunc("/", handlers.HelloAPI())
	mux.HandleFunc("GET /posts", handle.ListHandlerPost(ctx))
	mux.HandleFunc("POST /posts", handle.CreateHandlerPost(ctx))
}
