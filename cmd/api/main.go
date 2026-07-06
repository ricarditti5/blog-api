package main

import (
	"blog-api/internal/config"
	"blog-api/internal/services"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	HelloAPi := services.HelloAPI()

	//load config
	cfg, _ := config.LoadConfig()

	//init database
	db, _ := config.InitDB(cfg)

	mux.HandleFunc("/", HelloAPi)
	fmt.Println("Server Runnig...")
	http.ListenAndServe(":8080", mux)
}
