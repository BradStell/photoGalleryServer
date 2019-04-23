package images

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetSlideshowImagesHandler returns a list of images on the homepage slide show
func GetSlideshowImagesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slideshow := mux.Vars(r)["slideshow"]

		fmt.Fprintf(w, "slideshow: %s", slideshow)
	}
}

// GetAllImagesHandler returns a HandlerFunc for getting all images
func GetAllImagesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var images []Image

		images, err := GetAllImages(db)

		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(images)
	}
}

// CreateImageHandler returns a HandlerFunc for creating an image
func CreateImageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var image Image

		err := json.NewDecoder(r.Body).Decode(&image)

		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		imageDB, err := image.create(db)

		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(imageDB)
	}
}

// GetImageHandler returns a HandlerFunc for getting an image
func GetImageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageID := mux.Vars(r)["id"]

		image, err := GetImage(db, imageID)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(image)
	}
}

// UpdateImageHandler returns a HandlerFunc for updating an image
func UpdateImageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var image Image

		err := json.NewDecoder(r.Body).Decode(&image)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if string(image.ID) != mux.Vars(r)["id"] {
			fmt.Printf("Update error: image ID mismatch in update")
			http.Error(w, "Update error: image ID mismatch in update", http.StatusBadRequest)
			return
		}

		imageDB, err := image.update(db)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(imageDB)
	}
}

// DeleteImageHandler returns a HandlerFunc for deleting an image
func DeleteImageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageID := mux.Vars(r)["id"]

		err := DeleteImage(db, imageID)
		if err != nil {
			fmt.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}
}
