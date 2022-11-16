package songdto

type SongRequest struct{
	Title string `json :"title"`
	Image string `json :"image"`
	Year string `json : "year"`
	Song string `json :"song"`
	ArtisID int `json : "artis_id"`

}