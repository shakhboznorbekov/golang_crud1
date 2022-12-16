package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shakhboznorbekov/mytasks/golang_crud/api/docs"
	"github.com/shakhboznorbekov/mytasks/golang_crud/api/handler"
	"github.com/shakhboznorbekov/mytasks/golang_crud/storage"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(storage)

	r.POST("/film", handlerV1.CreateFilm)
	r.GET("/film/:id", handlerV1.GetFilmById)
	r.GET("/film", handlerV1.GetFilmList)
	r.PUT("/film/:id", handlerV1.UpdateFilm)
	r.DELETE("/film/:id", handlerV1.DeleteFilm)

	r.POST("/category", handlerV1.CreateFilm)
	r.GET("/category/:id", handlerV1.GetFilmById)
	r.GET("/category", handlerV1.GetFilmList)
	r.PUT("/category/:id", handlerV1.UpdateFilm)
	r.DELETE("/category/:id", handlerV1.DeleteFilm)

	r.POST("/actor", handlerV1.CreateFilm)
	r.GET("/actor/:id", handlerV1.GetFilmById)
	r.GET("/actor", handlerV1.GetFilmList)
	r.PUT("/actor/:id", handlerV1.UpdateFilm)
	r.DELETE("/actor/:id", handlerV1.DeleteFilm)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
