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

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (f *categoryRepo) Create(ctx context.Context, category *models.CreateCategory) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO film_category(
			category_id,
			film_id, 
			updated_at
		) VALUES ( $1, $2 , now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		category.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *categoryRepo) GetByPKey(ctx context.Context, pkey *models.CategoryPrimarKey) (*models.Category, error) {

	var (
		id         sql.NullString
		categoryId sql.NullString
		updatedAt  sql.NullString
	)

	query := `
		SELECT
			category_id,
			film_id, 
			updated_at
		FROM
			film_category
		WHERE category_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&categoryId,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:        id.String,
		Name:      categoryId.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (f *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	var (
		resp   = models.GetListCategoryResponse{}
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
			category_id,
			film_id, 
			updated_at
		FROM
			film_category
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id          sql.NullString
			category_id sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&category_id,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Categorys = append(resp.Categorys, &models.Category{
			Id:        id.String,
			Name:      category_id.String,
			UpdatedAt: updatedAt.String,
		})

	}

	return &resp, err
}

func (f *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			film_category
		SET
			category_id = :category_id,
			film_id = :film_id, 
			updated_at = now()
		WHERE category_id = :category_id
	`

	params = map[string]interface{}{
		"category_id": req.Name,
		"name":        req.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *categoryRepo) Delete(ctx context.Context, req *models.CategoryPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM film_category WHERE category_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
