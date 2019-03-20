package routes

import (
	"database/sql"
	"go-server-test/galleries"
	"go-server-test/images"
	"go-server-test/middleware"

	"github.com/gorilla/mux"
)

func Init(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Routes
	imagesRouter := r.PathPrefix("/images").Subrouter()
	galleriesRouter := r.PathPrefix("/galleries").Subrouter()

	// Image routes
	imagesRouter.HandleFunc("/", images.GetAllHandler(db)).Methods("GET")        // GET
	imagesRouter.HandleFunc("/", images.CreateHandler(db)).Methods("POST")       // POST
	imagesRouter.HandleFunc("/{id}", images.GetHandler(db)).Methods("GET")       // GET
	imagesRouter.HandleFunc("/{id}", images.UpdateHandler(db)).Methods("PUT")    // PUT
	imagesRouter.HandleFunc("/{id}", images.DeleteHandler(db)).Methods("DELETE") // DELETE

	// Gallery routes
	galleriesRouter.HandleFunc("/", galleries.GetAllHandler(db)).Methods("GET")        // GET
	galleriesRouter.HandleFunc("/", galleries.CreateHandler(db)).Methods("POST")       // POST
	galleriesRouter.HandleFunc("/{id}", galleries.GetHandler(db)).Methods("GET")       // GET
	galleriesRouter.HandleFunc("/{id}", galleries.UpdateHandler(db)).Methods("PUT")    // PUT
	galleriesRouter.HandleFunc("/{id}", galleries.DeleteHandler(db)).Methods("DELETE") // DELETE
	galleriesRouter.HandleFunc("/{id}/images", galleries.GetImages(db)).Methods("GET") // GET

	r.Use(middleware.LoggingMiddleware)

	return r
}
