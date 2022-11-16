package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type ArtisRepository interface {
	CreateArtis(artis models.Artis) (models.Artis, error)
	FindArtis() ([]models.Artis, error)
	GetArtisById(ID int) (models.Artis, error)
}
func RepositoryArtis(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) CreateArtis(artis models.Artis) (models.Artis,error){
err := r.db.Create(&artis).Error

return artis,err
}
func (r *repository) FindArtis() ([]models.Artis,error){
	var artis []models.Artis
	err := r.db.Find(&artis).Error

	return artis,err
}
func (r *repository) GetArtisById(ID int) (models.Artis,error){
var artis models.Artis
err := r.db.First(&artis,ID).Error

return artis,err
}