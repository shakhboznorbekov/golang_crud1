package storage_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/config"
	"github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/storage/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	filmRepo     *postgres.FilmRepo
	actorRepo    *postgres.ActorRepo
	categoryRepo *postgres.CategoryRepo
)

func TestMain(m *testing.M) {

	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	filmRepo = postgres.NewFilmRepo(pool)
	actorRepo = postgres.NewActorRepo(pool)
	categoryRepo = postgres.NewCategoryRepo(pool)

	os.Exit(m.Run())
}
