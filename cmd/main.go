package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/api"
	"github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/config"
	"github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/storage/postgres"
)

func main() {

	cfg := config.Load()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	storage, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.CloseDB()

	api.SetUpApi(r, storage)

	log.Printf("Listening port %v...\n", cfg.HTTPPort)
	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
