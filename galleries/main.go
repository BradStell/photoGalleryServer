package galleries

import (
	"database/sql"
	"fmt"
	"go-server-test/images"
	"net/http"

	"github.com/gorilla/mux"
)

type Gallery struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	CoverImageID int           `json:"coverImageId"`
	CoverImage   *images.Image `json:"coverImage"`
	Created      int           `json:"created"`
	Modified     int           `json:"modified"`
	Order        int           `json:"order"`
}

func (i *Gallery) save() string {
	return "New Gallery"
}

func GetAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /galleries")
	}
}

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "POST /galleries")
	}
}

func GetHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /galleries/"+mux.Vars(r)["id"])
	}
}

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "PUT /galleries/"+mux.Vars(r)["id"])
	}
}

func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DELETE /galleries/"+mux.Vars(r)["id"])
	}
}

func GetImages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GET /galleries/"+mux.Vars(r)["id"]+"/images")
	}
}
