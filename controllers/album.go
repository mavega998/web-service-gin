package controllers

import (
	"example/web-service-gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /albums
func FindAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)
	c.JSON(http.StatusOK, gin.H{"data": albums})
}

// POST /albums
func CreateAlbum(c *gin.Context) {
	var albumNew models.AlbumInput
	if err := c.ShouldBindJSON(&albumNew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album := models.Album{Title: albumNew.Title, Artist: albumNew.Artist, Price: albumNew.Price}
	id := models.DB.Create(&album)
	fmt.Print(id.Error)
	c.JSON(http.StatusCreated, gin.H{"data": album})
}

// GET /albums/:id
func FindAlbumById(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	if err := models.DB.Where("id = ?", id).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	if err := models.DB.Where("id = ?", id).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found!"})
		return
	}

	var albumUpd models.AlbumUpdate
	if err := c.ShouldBindJSON(&albumUpd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&album).Updates(albumUpd)
	c.JSON(http.StatusOK, gin.H{"data": album})
}
