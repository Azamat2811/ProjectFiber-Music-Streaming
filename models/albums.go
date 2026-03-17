package models

type Albums struct {
	ID            int    `json:"id"`
	ArtistId      int    `json:"artist_id"`
	Name          string `json:"name"`
	YearOfRelease int    `json:"year_of_release"`
}
