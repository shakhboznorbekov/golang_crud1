package storage

import (
	"context"

	"github.com/shakhboznorbekov/mytasks/golang_crud/models"
)

type StorageI interface {
	CloseDB()
	Film() FilmRepoI
	Actor() ActorRepo
	Category() CategoryRepo
}

type FilmRepoI interface {
	Create(ctx context.Context, req *models.CreateFilm) (string, error)
	GetByPKey(ctx context.Context, req *models.FilmPrimarKey) (*models.Film, error)
	GetList(ctx context.Context, req *models.GetListFilmRequest) (*models.GetListFilmResponse, error)
	Update(ctx context.Context, req *models.UpdateFilm) (int64, error)
	Delete(ctx context.Context, req *models.FilmPrimarKey) error
}

type ActorRepo interface {
	Create(ctx context.Context, req *models.CreateActor) (string, error)
	GetByPKey(ctx context.Context, req *models.ActorPrimarKey) (*models.Actor, error)
	GetList(ctx context.Context, req *models.GetListActorRequest) (*models.GetListActorResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateActor) (int64, error)
	Delete(ctx context.Context, req *models.ActorPrimarKey) error
}

type CategoryRepo interface {
	Create(ctx context.Context, req *models.CreateCategory) (string, error)
	GetByPKey(ctx context.Context, req *models.CategoryPrimarKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.CategoryPrimarKey) error
}