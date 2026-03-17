package services

import (
	"database/sql"
	"errors"
	"projectfiber/db"
	"projectfiber/models"
)

func CreateTrack(track models.Tracks) (models.Tracks, error) {
	err := db.DB.QueryRow("INSERT INTO tracks (album_id, name, play_count) VALUES ($1, $2, $3) RETURNING id",
		track.Album_id,
		track.Name,
		track.Play_count,
	).Scan(&track.ID)

	if err != nil {
		return models.Tracks{}, err
	}
	return track, nil
}

func GetTrackByID(id int) (models.Tracks, error) {
	var track models.Tracks
	err := db.DB.QueryRow("SELECT id, album_id, name, play_count FROM tracks WHERE id = $1",
		id,
	).Scan(&track.ID, &track.Album_id, &track.Name, &track.Play_count)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Tracks{}, errors.New("track not found")
		}
		return models.Tracks{}, err
	}
	return track, nil
}

func GetAllTracks() ([]models.Tracks, error) {
	rows, err := db.DB.Query("SELECT id, album_id, name, play_count FROM tracks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []models.Tracks
	for rows.Next() {
		var track models.Tracks
		if err := rows.Scan(&track.ID, &track.Album_id, &track.Name, &track.Play_count); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func DeleteTrack(id int) error {
	result, err := db.DB.Exec(
		"DELETE FROM tracks WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("track not found")
	}

	return nil
}

func UpdateTrack(id int, data map[string]interface{}) (models.Tracks, error) {
	track, err := GetTrackByID(id)
	if err != nil {
		return models.Tracks{}, err
	}
	if album_id, ok := data["album_id"].(float64); ok {
		track.Album_id = int(album_id)
	}
	if name, ok := data["name"].(string); ok {
		track.Name = name
	}
	if play_count, ok := data["play_count"].(float64); ok {
		track.Play_count = int(play_count)
	}

	_, err = db.DB.Exec(

		"UPDATE tracks SET album_id = $1, name = $2, play_count = $3 WHERE id = $4",
		track.Album_id,
		track.Name,
		track.Play_count,
		id,
	)
	if err != nil {
		return models.Tracks{}, err
	}
	return track, nil
}

func FullUpdateTracks(id int, updated models.Tracks) (models.Tracks, error) {
	result, err := db.DB.Exec(
		"UPDATE tracks SET album_id = $1, name = $2, play_count = $3 WHERE id = $4",
		updated.Album_id,
		updated.Name,
		updated.Play_count,
		id,
	)

	if err != nil {
		return models.Tracks{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return models.Tracks{}, err
	}
	if rows == 0 {
		return models.Tracks{}, errors.New("track not found")
	}

	updated.ID = id
	return updated, nil
}

func GetTracksByAlbum(albumID int) ([]models.Tracks, error) {
	rows, err := db.DB.Query(
		"SELECT id, album_id, name, play_count FROM tracks WHERE album_id = $1",
		albumID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []models.Tracks
	for rows.Next() {
		var track models.Tracks
		if err := rows.Scan(&track.ID, &track.Album_id, &track.Name, &track.Play_count); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}
