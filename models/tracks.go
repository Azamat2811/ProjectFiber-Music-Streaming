package models

type Tracks struct {
	ID         int    `json:"id"`
	Album_id   int    `json:"album_id"`
	Name       string `json:"name"`
	Play_count int    `json:"play_count"`
}
