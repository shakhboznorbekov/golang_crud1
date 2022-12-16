package models

type ActorPrimarKey struct {
	Id string `json:"actor_id"`
}

type CreateActor struct {
	FilmId string `json:"film_id"`
}

type Actor struct {
	Id        string `json:"actor_id"`
	FilmId    string `json:"film_id"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateActor struct {
	Id     string `json:"actor_id"`
	FilmId string `json:"film_id"`
}

type GetListActorRequest struct {
	Limit  int32
	Offset int32
}

type GetListActorResponse struct {
	Count  int32    `json:"count"`
	Actors []*Actor `json:"actors"`
}
