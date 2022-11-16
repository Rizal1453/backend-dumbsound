package models



type Artis struct {
	ID        int       `json:"id"`
	Name string `json:"name"`
	Old 	 string    `json:"old"`
	Role 	string    `json:"role"`
	StartCareer string `json:"startcareer"`
	Songs []Song `json:"songs"`
}
type ArtistProfile struct {
	ID          int    `json:"id"  gorm:"primary_key:auto_increment"`
	Name        string `json:"name"`
	Old         int    `json:"old"`
	Role        string `json:"role"`
	StartCareer int    `json:"start_career"`
}

func (ArtistProfile) TableName() string {
	return "artis"
}

	