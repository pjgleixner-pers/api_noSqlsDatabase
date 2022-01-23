package controller

import (
	"example/web-service-gin/model"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", model.GetAlbums)
	router.GET("/albums/:id", model.GetAlbumByID)
	router.POST("/albums", model.PostAlbums)

	router.Run("localhost:8080")

	return router
}
