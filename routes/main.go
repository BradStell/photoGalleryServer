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
	imagesRouter.HandleFunc("/", images.GetAllImagesHandler(db)).Methods("GET")                                          // GET
	imagesRouter.HandleFunc("", images.GetSlideshowImagesHandler(db)).Methods("GET").Queries("slideshow", "{slideshow}") // GET
	imagesRouter.HandleFunc("/", images.CreateImageHandler(db)).Methods("POST")                                          // POST
	imagesRouter.HandleFunc("/{id}", images.GetImageHandler(db)).Methods("GET")                                          // GET
	imagesRouter.HandleFunc("/{id}", images.UpdateImageHandler(db)).Methods("PUT")                                       // PUT
	imagesRouter.HandleFunc("/{id}", images.DeleteImageHandler(db)).Methods("DELETE")                                    // DELETE

	// Gallery routes
	galleriesRouter.HandleFunc("/", galleries.GetAllGalleriesHandler(db)).Methods("GET")      // GET
	galleriesRouter.HandleFunc("/", galleries.CreateGalleryHandler(db)).Methods("POST")       // POST
	galleriesRouter.HandleFunc("/{id}", galleries.GetGalleryHandler(db)).Methods("GET")       // GET
	galleriesRouter.HandleFunc("/{id}", galleries.UpdateGalleryHandler(db)).Methods("PUT")    // PUT
	galleriesRouter.HandleFunc("/{id}", galleries.DeleteGalleryHandler(db)).Methods("DELETE") // DELETE
	galleriesRouter.HandleFunc("/{id}/images", galleries.GetGalleryImages(db)).Methods("GET") // GET

	// Mount middleware
	r.Use(middleware.LoggingMiddleware)

	return r
}

// Routes Needed
/*
	Home Page
		- Get all slideshow images
		- Get about me image (hard coded resource ??)
		- List of gallery names for menu dropdown

	Gallery Pages
	 	- Get gallery information (name, date, etc...)
		- Get gallery cover image
		- Get all images for gallery (low resolution **)
		- Get full size image for each image (maybe on initial load - maybe lazy loaded)
*/
