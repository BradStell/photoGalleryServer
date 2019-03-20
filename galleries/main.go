package galleries

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Gallery struct {
	ID   string
	Name string
}

func (i *Gallery) save() string {
	return "New Gallery"
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET /galleries")
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST /galleries")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET /galleries/"+mux.Vars(r)["id"])
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT /galleries/"+mux.Vars(r)["id"])
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE /galleries/"+mux.Vars(r)["id"])
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET /galleries/"+mux.Vars(r)["id"]+"/images")
}
