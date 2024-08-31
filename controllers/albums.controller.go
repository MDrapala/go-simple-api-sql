package controllers

import (
	"fmt"
	"net/http"

	DB "api-sql/config"
	Album "api-sql/types"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []Album.Album

	rows, err := DB.DB.Query("SELECT * FROM album")
	if err != nil {
		fmt.Errorf("Error fetching albums: %v", err)
		return
	}

	for rows.Next() {
		var album Album.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			fmt.Errorf("Error scanning album row: %v", err)
			continue
		}
		albums = append(albums, album)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	var album Album.Album

	err := DB.DB.QueryRow("SELECT * FROM album WHERE id = ?", id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		fmt.Errorf("Error fetching album: %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func PostAlbum(c *gin.Context) {
	var album Album.Album
	c.BindJSON(&album)

	_, err := DB.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		fmt.Errorf("Error inserting album: %v", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, album)
}

func PutAlbum(c *gin.Context) {
	id := c.Param("id")
	var album Album.Album
	c.BindJSON(&album)

	_, err := DB.DB.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?", album.Title, album.Artist, album.Price, id)
	if err != nil {
		fmt.Errorf("Error updating album: %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	_, err := DB.DB.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		fmt.Errorf("Error deleting album: %v", err)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
