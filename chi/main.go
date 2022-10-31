package main

import (
	"fmt"
	"net/http"

	demo "github.com/fgm/go__web_demo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// main start the program with os signals support.
func main() {
	r := chi.NewRouter()
	r.Get(demo.RouteAlbums, getAlbums())
	r.Get(demo.RouteSingleAlbum, getAlbumByID())
	r.Post(demo.RouteAlbums, postAlbums())
	
	srv := new(http.Server)
	srv.Addr=":8080" 
	srv.Handler = r
	srv.ListenAndServe()
}

// getAlbums provide a request handler to getAlbums all albums.
func getAlbums() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, demo.Albums)
	}
}

// getAlbumByID provide a request handler to getAlbumByID one album.
func getAlbumByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		for _, album := range demo.Albums {
			if album.ID == id {
				render.JSON(w, r, album)
				return
			}
		}
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, ErrResponse{
			Message: "album not found",
		})
	}
}

// ErrResponse provide general purpose JSON response.
type ErrResponse struct {
	Message string `json:"message,omitempty"` // user-level status message
}

// Render adapter function for Renderer.
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// AlbumRequest decouple request from data model.
type AlbumRequest struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Bind adapter function.
func (album *AlbumRequest) Bind(r *http.Request) error {
	if album.ID == "" {
		return fmt.Errorf("invalid")
	}
	return nil
}

// postAlbums provide a request handler for album creation.
func postAlbums() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		album := new(AlbumRequest)

		// Bind the received JSON to album
		if err := render.Bind(r, album); err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, nil)
			return
		}

		// Add the new Album to the slice
		demo.Albums = append(demo.Albums, demo.Album{
			ID:     album.ID,
			Title:  album.Title,
			Artist: album.Artist,
			Price:  album.Price,
		})
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, album)
	}
}
