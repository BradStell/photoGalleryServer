package images

import (
	"database/sql"
	"errors"
	"time"

	// Postgresql driver
	_ "github.com/lib/pq"
)

// Image data structure for storing image meta data information
type Image struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Path          string    `json:"path"`
	IsOnSlideshow bool      `json:"isOnSlideshow"`
	Created       time.Time `json:"created"`
	Modified      time.Time `json:"modified"`
	HighResID     *int      `json:"highResID"`
	HighRes       *Image    `json:"highResMe"`
}

func (i *Image) create(db *sql.DB) (*Image, error) {
	rows, err := db.Query(`
	INSERT INTO ImageMeta (title, path, high_res_id)
	VALUES ($1, $2, $3)
	RETURNING id, title, path, created, modified, high_res_id
	`, i.Title, i.Path, i.HighResID)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var image Image
	if rows.Next() {
		err = rows.Scan(&image.ID, &image.Title, &image.Path, &image.Created, &image.Modified, &image.HighResID)

		if err != nil {
			return nil, err
		}
	}

	return &image, nil
}

func (i *Image) update(db *sql.DB) (*Image, error) {
	rows, err := db.Query(`
		UPDATE ImageMeta 
		SET title = $1, path = $2, high_res_id = $3
		WHERE id = $4
		RETURNING id, title, path, created, modified, high_res_id
	`, i.Title, i.Path, i.HighResID, i.ID)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var image Image
	if rows.Next() {
		err = rows.Scan(&image.ID, &image.Title, &image.Path, &image.Created, &image.Modified, &image.HighResID)

		if err != nil {
			return nil, err
		}
	}

	return &image, nil
}

func (i *Image) delete(db *sql.DB) error {
	res, err := db.Exec("DELETE FROM ImageMeta WHERE id = $1", i.ID)
	if err != nil {
		return err
	}

	success, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if success != 1 {
		return errors.New("Image not deleted")
	}

	return nil
}

// DeleteImage attempts to deletes an image from the DB
func DeleteImage(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM ImageMeta WHERE id = $1", id)
	if err != nil {
		return err
	}

	success, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if success != 1 {
		return errors.New("Image not deleted")
	}

	return nil
}

// GetImage attempts to get an image from the DB by the specified id
func GetImage(db *sql.DB, id string) (*Image, error) {
	rows, err := db.Query(`
		SELECT id, title, path, created, modified, high_res_id
		FROM ImageMeta
		WHERE id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	var image Image
	if rows.Next() {
		err := rows.Scan(&image.ID, &image.Title, &image.Path, &image.Created, &image.Modified, &image.HighResID)

		if err != nil {
			return nil, err
		}
	}

	return &image, nil
}

// GetAllImages attempts to get all image meta records from the DB
func GetAllImages(db *sql.DB) ([]Image, error) {
	var images []Image

	rows, err := db.Query(`
		SELECT id, title, path, created, modified, high_res_id
		FROM ImageMeta
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var image Image
		err = rows.Scan(&image.ID, &image.Title, &image.Path, &image.Created, &image.Modified, &image.HighResID)

		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}
