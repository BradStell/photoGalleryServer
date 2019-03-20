package images

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(images)
	}
}

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var image Image
		err := decoder.Decode(&image)
		if err != nil {
			panic(err)
		}
		imageDB := image.create(db)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(imageDB)
	}
}

func GetHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := mux.Vars(r)
		id := queryParams["id"]

		image := Get(db, id)

		json.NewEncoder(w).Encode(image)
	}
}

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var image Image
		image.update()
		json.NewEncoder(w).Encode(image)
	}
}

func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var image Image
		image.delete()
		fmt.Fprintf(w, "DELETE /images/"+mux.Vars(r)["id"])
	}
}
