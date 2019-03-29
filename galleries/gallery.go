package galleries

import (
	"database/sql"
	"errors"
	"go-server-test/images"
	"time"

	// Postgresql driver
	_ "github.com/lib/pq"
)

// Gallery meta data data structure
type Gallery struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	CoverImageID *int          `json:"coverImageId"`
	CoverImage   *images.Image `json:"coverImage"`
	Created      time.Time     `json:"created"`
	Modified     time.Time     `json:"modified"`
	Priorety     int           `json:"priority"`
}

func (i *Gallery) create(db *sql.DB) (*Gallery, error) {
	rows, err := db.Query(`
		INSERT INTO GalleryMeta (name, cover_image_id, priorety)
		VALUES $1, $2, $3
		RETURNING id, name, cover_image_id, created, modified, priorety
	`, i.Name, i.CoverImageID, i.Priorety)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var gallery Gallery
	if rows.Next() {
		err = rows.Scan(&gallery.ID, &gallery.Name, &gallery.CoverImageID, &gallery.Created, &gallery.Modified, &gallery.Priorety)

		if err != nil {
			return nil, err
		}
	}

	return &gallery, nil
}

func (i *Gallery) update(db *sql.DB) (*Gallery, error) {
	rows, err := db.Query(`
		UPDATE GalleryMeta
		SET name = $1, cover_image_id = $2, priorety = $3
		WHERE id = $4
		RETURNING id, name, cover_image_id, created, modified, priorety
	`, i.Name, i.CoverImageID, i.Priorety, i.ID)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var gallery Gallery
	if rows.Next() {
		err = rows.Scan(&gallery.ID, &gallery.Name, &gallery.CoverImageID, &gallery.Created, &gallery.Modified, &gallery.Priorety)

		if err != nil {
			return nil, err
		}
	}

	return &gallery, nil
}

func (i *Gallery) delete(db *sql.DB) error {
	res, err := db.Exec("DELETE FROM GalleryMeta WHERE id = $1", i.ID)
	if err != nil {
		return err
	}

	success, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if success != 1 {
		return errors.New("Gallery not deleted")
	}

	return nil
}

// DeleteGallery deletes the specified record from the DB
func DeleteGallery(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM GalleryMeta WHERE id = $1", id)
	if err != nil {
		return err
	}

	success, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if success != 1 {
		return errors.New("Gallery not deleted")
	}

	return nil
}

// GetGallery attempts to retreive the specified gallery record from the DB, specified by the id
func GetGallery(db *sql.DB, id string) (*Gallery, error) {
	rows, err := db.Query(`
		SELECT id, name, cover_image_id, created, modified, priorety
		FROM GalleryMeta
		WHERE id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	var gallery Gallery
	if rows.Next() {
		err = rows.Scan(&gallery.ID, &gallery.Name, &gallery.CoverImageID, &gallery.Created, &gallery.Modified, &gallery.Priorety)

		if err != nil {
			return nil, err
		}
	}

	return &gallery, nil
}

// GetAllGalleries attempts to retreive all galleries from the gallery table
func GetAllGalleries(db *sql.DB) ([]Gallery, error) {
	var galleries []Gallery

	rows, err := db.Query(`
		SELECT id, name, cover_image_id, created, modified, priorety
		FROM GalleryMeta
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var gallery Gallery
		err = rows.Scan(&gallery.ID, &gallery.Name, &gallery.CoverImageID, &gallery.Created, &gallery.Modified, &gallery.Priorety)

		if err != nil {
			return nil, err
		}

		galleries = append(galleries, gallery)
	}

	return galleries, nil
}
