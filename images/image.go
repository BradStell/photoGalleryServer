package images

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Image struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	Created   int    `json:"created"`
	Modified  int    `json:"modified"`
	HighResID int    `json:"highResID"`
	HighRes   *Image `json:"highResMe"`
}

func (i *Image) create(db *sql.DB) *Image {
	var image Image

	db.QueryRow(`
		INSERT INTO ImageMeta (title, path, high_res_id)
		VALUES ($1, $2, $3)
		RETURNING *	
	`, i.Title, i.Path, i.HighResID).Scan(&image)

	return &image
}

func (i *Image) update() *Image {
	return nil
}

func (i *Image) delete() *Image {
	return nil
}

func (i *Image) softDelete() *Image {
	return nil
}

func Get(db *sql.DB, id string) *Image {
	var image Image

	db.QueryRow("SELECT * FROM ImageMeta WHERE id = $1", id).Scan(&image)

	return &image
}

func GetAll() []*Image {
	return nil
}
