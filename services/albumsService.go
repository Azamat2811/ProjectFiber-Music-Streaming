package services

import (
	"database/sql"
	"errors"
	"projectfiber/db"
	"projectfiber/models"
)

func CreateAlbum(album models.Albums) (models.Albums, error) {

	err := db.DB.QueryRow("INSERT INTO albums (artist_id, name, year_of_release) VALUES ($1, $2, $3) RETURNING id",
		album.ArtistId,
		album.Name,
		album.YearOfRelease,
	).Scan(&album.ID)

	if err != nil {
		return models.Albums{}, err
	}
	return album, nil
}

func GetAlbumByID(id int) (models.Albums, error) {
	var album models.Albums
	err := db.DB.QueryRow("SELECT id, artist_id, name, year_of_release FROM albums WHERE id = $1",
		id,
	).Scan(&album.ID, &album.ArtistId, &album.Name, &album.YearOfRelease)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Albums{}, errors.New("album not found")
		}
		return models.Albums{}, err
	}
	return album, nil
}

func GetAllAlbums() ([]models.Albums, error) {
	rows, err := db.DB.Query("SELECT id, artist_id, name, year_of_release FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []models.Albums
	for rows.Next() {
		var album models.Albums
		if err := rows.Scan(&album.ID, &album.ArtistId, &album.Name, &album.YearOfRelease); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func DeleteAlbum(id int) error {

	result, err := db.DB.Exec(
		"DELETE FROM albums WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("album not found")
	}

	return nil
}

func UpdateAlbum(id int, data map[string]interface{}) (models.Albums, error) {
	album, err := GetAlbumByID(id)
	if err != nil {
		return models.Albums{}, err
	}

	if name, ok := data["name"].(string); ok {
		album.Name = name
	}
	if year_of_release, ok := data["year_of_release"].(float64); ok {
		album.YearOfRelease = int(year_of_release)
	}
	if artist_id, ok := data["artist_id"].(float64); ok {
		album.ArtistId = int(artist_id)
	}

	_, err = db.DB.Exec(

		"UPDATE albums SET artist_id = $1, name = $2, year_of_release = $3 WHERE id = $4",
		album.ArtistId,
		album.Name,
		album.YearOfRelease,
		id,
	)
	if err != nil {
		return models.Albums{}, err
	}
	return album, nil
}

func FullUpdateAlbums(id int, updated models.Albums) (models.Albums, error) {
	result, err := db.DB.Exec(
		"UPDATE albums SET artist_id = $1, name = $2, year_of_release = $3 WHERE id = $4",
		updated.ArtistId,
		updated.Name,
		updated.YearOfRelease,
		id,
	)

	if err != nil {
		return models.Albums{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return models.Albums{}, err
	}
	if rows == 0 {
		return models.Albums{}, errors.New("album not found")
	}

	updated.ID = id
	return updated, nil
}
