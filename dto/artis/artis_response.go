package artisdto

type ArtisResponse struct {
	ID        int       `json:"id"`
	Name string `gorm:"type: varchar(255)" json:"name"`
	Old 	 string    `json:"old"`
	Categori string    `json:"categori"`
	StartCareer string `json:"startcareer"`
}