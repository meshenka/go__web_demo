package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	demo "github.com/fgm/go__web_demo"
)

// getAlbums responds with the list of all albums as JSON.a
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, demo.Albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum demo.Album

	// Bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new Album to the slice
	demo.Albums = append(demo.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range demo.Albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}

func main() {
	gin.SetMode(gin.DebugMode) // or gin.TestMode or gin.ReleaseMode
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Route
	router.GET(demo.RouteAlbums, getAlbums)
	router.GET(demo.RouteSingleAlbum, getAlbumByID)
	router.POST(demo.RouteAlbums, postAlbums)

	// Handle
	router.Run("localhost:8081")
}
