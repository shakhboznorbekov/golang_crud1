package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/models"
	"github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/pkg/helper"
)

type actorRepo struct {
	db *pgxpool.Pool
}

func NewActorRepo(db *pgxpool.Pool) *actorRepo {
	return &actorRepo{
		db: db,
	}
}

func (f *actorRepo) Create(ctx context.Context, actor *models.CreateActor) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO film_actor(
			actor_id,
			film_id, 
			updated_at
		) VALUES ( $1, $2 , now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		actor.First_name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *actorRepo) GetByPKey(ctx context.Context, pkey *models.ActorPrimarKey) (*models.Actor, error) {

	var (
		id        sql.NullString
		filmId    sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			actor_id,
			film_id, 
			updated_at
		FROM
			film_actor
		WHERE actor_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&filmId,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Actor{
		Id:         id.String,
		First_name: filmId.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (f *actorRepo) GetList(ctx context.Context, req *models.GetListActorRequest) (*models.GetListActorResponse, error) {

	var (
		resp   = models.GetListActorResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 2"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			actor_id,
			film_id, 
			updated_at
		FROM
			film_actor
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id        sql.NullString
			actor_id  sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&actor_id,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Actors = append(resp.Actors, &models.Actor{
			Id:         id.String,
			First_name: actor_id.String,
			UpdatedAt:  updatedAt.String,
		})

	}

	return &resp, err
}

func (f *actorRepo) Update(ctx context.Context, req *models.UpdateActor) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			film_actor
		SET
			actor_id = :actor_id,
			film_id = :film_id, 
			updated_at = now()
		WHERE actor_id = :actor_id
	`

	params = map[string]interface{}{
		"actor_id":   req.First_name,
		"first_name": req.First_name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *actorRepo) Delete(ctx context.Context, req *models.ActorPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM film_actor WHERE actor_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
