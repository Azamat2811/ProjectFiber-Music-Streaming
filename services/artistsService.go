package services

import (
	"database/sql"
	"errors"
	"projectfiber/db"
	"projectfiber/models"
)

func CreateArtist(artist models.Artists) (models.Artists, error) {
	err := db.DB.QueryRow("INSERT INTO artists (name, genre, monthly_listeners) VALUES ($1, $2, $3) RETURNING id",
		artist.Name,
		artist.Genre,
		artist.Monthly_listeners,
	).Scan(&artist.ID)

	if err != nil {
		return models.Artists{}, err
	}
	return artist, nil
}

func GetArtistByID(id int) (models.Artists, error) {
	var artist models.Artists
	err := db.DB.QueryRow("SELECT id, name, genre, monthly_listeners FROM artists WHERE id = $1",
		id,
	).Scan(&artist.ID, &artist.Name, &artist.Genre, &artist.Monthly_listeners)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Artists{}, errors.New("artist not found")
		}
		return models.Artists{}, err
	}
	return artist, nil
}

func GetAllArtists() ([]models.Artists, error) {
	rows, err := db.DB.Query("SELECT id, name, genre, monthly_listeners FROM artists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []models.Artists
	for rows.Next() {
		var artist models.Artists
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.Genre, &artist.Monthly_listeners); err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	return artists, nil
}

func DeleteArtist(id int) error {
	result, err := db.DB.Exec(
		"DELETE FROM artists WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("artist not found")
	}

	return nil
}

func UpdateArtist(id int, data map[string]interface{}) (models.Artists, error) {
	artist, err := GetArtistByID(id)
	if err != nil {
		return models.Artists{}, err
	}

	if name, ok := data["name"].(string); ok {
		artist.Name = name
	}
	if genre, ok := data["genre"].(string); ok {
		artist.Genre = genre
	}
	if monthly_listeners, ok := data["monthly_listeners"].(float64); ok {
		artist.Monthly_listeners = int(monthly_listeners)
	}

	_, err = db.DB.Exec(

		"UPDATE artists SET name = $1, genre = $2, monthly_listeners = $3 WHERE id = $4",
		artist.Name,
		artist.Genre,
		artist.Monthly_listeners,
		id,
	)
	if err != nil {
		return models.Artists{}, err
	}
	return artist, nil
}

func FullUpdateArtists(id int, updated models.Artists) (models.Artists, error) {
	result, err := db.DB.Exec(
		"UPDATE artists SET name = $1, genre = $2, monthly_listeners = $3 WHERE id = $4",
		updated.Name,
		updated.Genre,
		updated.Monthly_listeners,
		id,
	)

	if err != nil {
		return models.Artists{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return models.Artists{}, err
	}
	if rows == 0 {
		return models.Artists{}, errors.New("artist not found")
	}

	updated.ID = id
	return updated, nil
}

func GetAlbumsByArtist(artistID int) ([]models.Albums, error) {
	rows, err := db.DB.Query(
		"SELECT id, artist_id, name, year_of_release FROM albums WHERE artist_id = $1",
		artistID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums = []models.Albums{}
	for rows.Next() {
		var album models.Albums
		if err := rows.Scan(&album.ID, &album.ArtistId, &album.Name, &album.YearOfRelease); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}
