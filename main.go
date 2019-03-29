package main

import (
	"database/sql"
	"fmt"
	"go-server-test/routes"
	"log"
	"net/http"
	"os"
)

const (
	DB_CONN_STR string = "DB_CONN_STR"
)

func main() {
	// Initialize the DB
	db, err := initDbCon(os.Getenv(DB_CONN_STR))
	if err != nil {
		log.Fatal(err)
	}

	// Seed the DB schema
	if err := seedDB(db); err != nil {
		log.Fatal(err)
	}

	// Create and initialize API routes
	router := routes.Init(db)

	// Set root route for returning API information
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Version: v%s", "1.0.0")
	})

	port := os.Getenv("port")
	if port == "" {
		port = "5000"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func initDbCon(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	var now string
	rows := db.QueryRow("SELECT NOW()")

	if err := rows.Scan(&now); err != nil {
		return nil, err
	}

	fmt.Printf("Current time is %s\n", now)

	return db, nil
}

func seedDB(db *sql.DB) error {
	sql := `
		CREATE TABLE IF NOT EXISTS ImageMeta(
			id SERIAL PRIMARY KEY,
			title varchar(80),
			path varchar(100),
			created timestamptz DEFAULT now(),
			modified timestamptz DEFAULT now(),
			high_res_id integer,
			FOREIGN KEY(high_res_id) REFERENCES ImageMeta(id)
		);

		CREATE TABLE IF NOT EXISTS GalleryMeta(
			id SERIAL PRIMARY KEY,
			name varchar(80),
			cover_image_id integer,
			created timestamptz DEFAULT now(),
			modified timestamptz DEFAULT now(),
			priorety integer DEFAULT 0,
			FOREIGN KEY(cover_image_id) REFERENCES ImageMeta(id)
		);

		CREATE TABLE IF NOT EXISTS ImageGalleryMeta(
			image_id integer NOT NULL,
			gallery_id integer NOT NULL,
			created timestamptz DEFAULT now(),
			priorety integer DEFAULT 0,
			PRIMARY KEY (image_id, gallery_id) 
		);

		CREATE TABLE IF NOT EXISTS SlideShow(
			id SERIAL PRIMARY KEY,
			image_id integer,
			priorety integer DEFAULT 0,
			created timestamptz DEFAULT now(),
			FOREIGN KEY(image_id) REFERENCES ImageMeta(id)
		);
	`

	_, err := db.Exec(sql)

	if err != nil {
		return err
	}

	fmt.Println("DB seeded")
	return nil
}
