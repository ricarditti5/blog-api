package main

import (
	"blog-api/internal/config"
	"blog-api/internal/handlers"
	"blog-api/internal/repository"
	"blog-api/internal/routes"
	"blog-api/internal/services"
	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error to load config:%v ", err)
		os.Exit(1)
	}

	pool, err := config.InitDB(cfg)
	if err != nil {
		fmt.Printf("Error to connect to database: %v", err)
		os.Exit(1)
	}
	repo := repository.NewPostRepo(pool)
	svc := services.NewPostService(repo)
	h := handlers.NewHandlePost(svc)

	mux := http.NewServeMux()
	ctx := context.Background()

	routes.RegisterRoutes(mux, ctx, h)

	fmt.Println("Server Runnig...")
	http.ListenAndServe(":8080", mux)
}
