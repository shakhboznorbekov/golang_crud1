package handler

import "github.com/shakhboznorbekov/mytasks/golang_crud/golang_crud1/storage"

type HandlerV1 struct {
	storage storage.StorageI
}

func NewHandlerV1(storage storage.StorageI) *HandlerV1 {
	return &HandlerV1{
		storage: storage,
	}
}
