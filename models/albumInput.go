package models

type AlbumInput struct {
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Price  uint   `json:"price" binding:"required"`
}

type AlbumUpdate struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  uint   `json:"price"`
}
