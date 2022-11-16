package models

type Song struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Year    string `json:"year"`
	Song    string `json:"song"`
	Image   string `json:"image"`
	ArtisID int    `json:"artis_id"`
	Artis  ArtistProfile  `json:"artis"`
}