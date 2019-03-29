package galleries

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllGalleriesHandler returns a HandlerFunc for getting all galleries
func GetAllGalleriesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var galleries []Gallery

		galleries, err := GetAllGalleries(db)

		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(galleries)
	}
}

// CreateGalleryHandler returns a HandlerFunc for creating a new gallery
func CreateGalleryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var gallery Gallery

		err := json.NewDecoder(r.Body).Decode(&gallery)

		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		galleryDB, err := gallery.create(db)

		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(galleryDB)
	}
}

// GetGalleryHandler returns a HandlerFunc for getting a gallery
func GetGalleryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		galleryID := mux.Vars(r)["id"]

		gallery, err := GetGallery(db, galleryID)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gallery)
	}
}

// UpdateGalleryHandler returns a HandlerFunc for updating a gallery
func UpdateGalleryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var gallery Gallery

		err := json.NewDecoder(r.Body).Decode(&gallery)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if string(gallery.ID) != mux.Vars(r)["id"] {
			fmt.Printf("Update error: galler ID mismatch in update")
			http.Error(w, "Update error: galler ID mismatch in update", http.StatusBadRequest)
			return
		}

		galleryDB, err := gallery.update(db)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(galleryDB)
	}
}

// DeleteGalleryHandler returns a HandlerFunc for deleting a gallery
func DeleteGalleryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		galleryID := mux.Vars(r)["id"]

		err := DeleteGallery(db, galleryID)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}
}

// GetGalleryImages returns a HandlerFunc for getting all images for a gallery
func GetGalleryImages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /galleries/"+mux.Vars(r)["id"]+"/images")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Not Implemented"))
	}
}
