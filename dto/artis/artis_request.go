package artisdto


type ArtisRequest struct {
	Name string `gorm:"type: varchar(255)" json:"name"`
	Old 	 string    `json:"old"`
	Role string    `json:"role"`
	StartCareer string `json:"startcareer"`
}
