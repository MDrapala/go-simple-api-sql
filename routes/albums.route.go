package routes

import (
	controllers "api-sql/controllers"

	"github.com/gin-gonic/gin"
)

func RouteAlbums(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.POST("/albums", controllers.PostAlbum)
	router.PUT("/albums/:id", controllers.PutAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
}
