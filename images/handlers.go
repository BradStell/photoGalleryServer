package images

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	images := GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var image Image
	err := decoder.Decode(&image)
	if err != nil {
		panic(err)
	}
	image.create()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(image)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	image := Get("1234")
	json.NewEncoder(w).Encode(image)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var image Image
	image.update()
	json.NewEncoder(w).Encode(image)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var image Image
	image.delete()
	fmt.Fprintf(w, "DELETE /images/"+mux.Vars(r)["id"])
}
