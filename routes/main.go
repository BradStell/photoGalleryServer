package routes

import (
	"go-server-test/galleries"
	"go-server-test/images"
	"go-server-test/middleware"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	// Routes
	imagesRouter := r.PathPrefix("/images").Subrouter()
	galleriesRouter := r.PathPrefix("/galleries").Subrouter()

	// Image routes
	imagesRouter.HandleFunc("/", images.GetAllHandler).Methods("GET")        // GET
	imagesRouter.HandleFunc("/", images.CreateHandler).Methods("POST")       // POST
	imagesRouter.HandleFunc("/{id}", images.GetHandler).Methods("GET")       // GET
	imagesRouter.HandleFunc("/{id}", images.UpdateHandler).Methods("PUT")    // PUT
	imagesRouter.HandleFunc("/{id}", images.DeleteHandler).Methods("DELETE") // DELETE

	// Gallery routes
	galleriesRouter.HandleFunc("/", galleries.GetAllHandler).Methods("GET")        // GET
	galleriesRouter.HandleFunc("/", galleries.CreateHandler).Methods("POST")       // POST
	galleriesRouter.HandleFunc("/{id}", galleries.GetHandler).Methods("GET")       // GET
	galleriesRouter.HandleFunc("/{id}", galleries.UpdateHandler).Methods("PUT")    // PUT
	galleriesRouter.HandleFunc("/{id}", galleries.DeleteHandler).Methods("DELETE") // DELETE
	galleriesRouter.HandleFunc("/{id}/images", galleries.GetImages).Methods("GET") // GET

	r.Use(middleware.LoggingMiddleware)

	return r
}
