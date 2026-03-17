package services

import (
	"database/sql"
	"errors"
	"projectfiber/db"
	"projectfiber/models"
)

func CreatePlaylist(playlists models.Playlists) (models.Playlists, error) {
	err := db.DB.QueryRow("INSERT INTO playlists (name) VALUES ($1) RETURNING id",
		playlists.Name,
	).Scan(&playlists.ID)

	if err != nil {
		return models.Playlists{}, err
	}
	return playlists, nil
}

func GetPlaylistByID(id int) (models.Playlists, error) {
	var playlist models.Playlists
	err := db.DB.QueryRow("SELECT id, name FROM playlists WHERE id = $1",
		id,
	).Scan(&playlist.ID, &playlist.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Playlists{}, errors.New("playlist not found")
		}
		return models.Playlists{}, err
	}
	return playlist, nil
}

func GetAllPlaylists() ([]models.Playlists, error) {
	rows, err := db.DB.Query("SELECT id, name FROM playlists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []models.Playlists
	for rows.Next() {
		var playlist models.Playlists
		if err := rows.Scan(&playlist.ID, &playlist.Name); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func DeletePlaylist(id int) error {
	result, err := db.DB.Exec(
		"DELETE FROM playlists WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("playlist not found")
	}

	return nil
}

func UpdatePlaylist(id int, data map[string]interface{}) (models.Playlists, error) {
	playlist, err := GetPlaylistByID(id)
	if err != nil {
		return models.Playlists{}, err
	}
	if name, ok := data["name"].(string); ok {
		playlist.Name = name
	}

	_, err = db.DB.Exec(

		"UPDATE playlists SET name = $1 WHERE id = $2",
		playlist.Name,
		id,
	)
	if err != nil {
		return models.Playlists{}, err
	}
	return playlist, nil
}

func FullUpdatePlaylists(id int, updated models.Playlists) (models.Playlists, error) {
	result, err := db.DB.Exec(
		"UPDATE playlists SET name = $1 WHERE id = $2",
		updated.Name,
		id,
	)

	if err != nil {
		return models.Playlists{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return models.Playlists{}, err
	}
	if rows == 0 {
		return models.Playlists{}, errors.New("playlist not found")
	}

	updated.ID = id
	return updated, nil
}

