package models

type CategoryPrimarKey struct {
	Id string `json:"category_id"`
}

type CreateCategory struct {
	FilmId string `json:"film_id"`
}

type Category struct {
	Id        string `json:"category_id"`
	FilmId    string `json:"film_id"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateCategory struct {
	Id     string `json:"category_id"`
	FilmId string `json:"film_id"`
}

type GetListCategoryRequest struct {
	Limit  int32
	Offset int32
}

type GetListCategoryResponse struct {
	Count      int32       `json:"count"`
	Categories []*Category `json:"categories"`
}
