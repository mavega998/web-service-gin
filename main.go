package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome")
}

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	router.GET("/", welcome)
	router.GET("/albums", controllers.FindAlbums)
	router.GET("/albums/:id", controllers.FindAlbumById)
	router.POST("/albums", controllers.CreateAlbum)
	router.PATCH("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
	router.Run()
}
