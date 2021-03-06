package model

import (
	"example/web-service-gin/data"
	"example/web-service-gin/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range data.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	//var newAlbum controller.
	var newAlbum views.Album
	err := c.BindJSON(&newAlbum)
	// Call BindJSON to bind the received JSON to
	// newAlbum.

	if err != nil {
		return
	}

	// Add the new album to the slice.
	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/*func PostAlbums(c *gin.Context) {
	//var newAlbum controller.
	var newAlbum views.Album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	for i := range newAlbum {
		err := c.BindJSON(&newAlbum[i])
		if err != nil {
			return
		}
		data.Albums = append(data.Albums, newAlbum[i])
		c.IndentedJSON(http.StatusCreated, newAlbum[i])
	}
}*/

func DeleteAlbums(c *gin.Context) {
	id := c.Param("id")

	for i, a := range data.Albums {
		if a.ID == id {
			data.Albums = append(data.Albums[:i], data.Albums[i+1:]...) //I dont get how it works
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
