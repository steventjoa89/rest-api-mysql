package entity

import "time"

// POST OBJECT
type Artist struct {
	ArtistId    string    `json:"artistId"`
	ArtistName  string    `json:"artistName"`
	AlbumName   string    `json:"albumName"`
	ImageURL    string    `json:"imageURL"`
	ReleaseDate time.Time `json:"releaseDate"`
	Price       float32   `json:"price"`
	SampleURL   string    `json:"sampleURL"`
}
