package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	demo "github.com/fgm/go__web_demo"
)

func indentedJSON(w http.ResponseWriter, status int, v any) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	enc.Encode(v)
}

// getAlbums responds with the list of all albums as JSON.a
func getAlbums(w http.ResponseWriter, r *http.Request) {
	indentedJSON(w, http.StatusOK, demo.Albums)
}

func postAlbums(w http.ResponseWriter, r *http.Request) {
	var newAlbum demo.Album

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed reading request body", http.StatusBadRequest)
		return
	}
	// Bind the received JSON to newAlbum
	if err := json.Unmarshal(body, &newAlbum); err != nil {
		http.Error(w, "failed decoding request body", http.StatusBadRequest)
		return
	}

	// Add the new Album to the slice
	demo.Albums = append(demo.Albums, newAlbum)
	indentedJSON(w, http.StatusCreated, newAlbum)
}

func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, album := range demo.Albums {
		if album.ID == id {
			indentedJSON(w, http.StatusOK, album)
			return
		}
	}
	indentedJSON(w, http.StatusNotFound, map[string]any{
		"message": "album not found",
	})
}

func main() {

	router := mux.NewRouter()
	// Set trusted proxies on handler, not on router.

	// Route
	router.HandleFunc(demo.RouteAlbums, getAlbums).Methods(http.MethodGet)
	router.HandleFunc(demo.RouteSingleAlbum, getAlbumByID).Methods(http.MethodGet)
	router.HandleFunc(demo.RouteAlbums, postAlbums).Methods(http.MethodPost)

	// Handle
	http.ListenAndServe("localhost:8080", handlers.ProxyHeaders(router))
}
