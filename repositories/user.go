package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User,error)
	FindUser()([]models.User,error)
	GetUserById(ID int) (models.User,error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) CreateUser(user models.User) (models.User,error){
	err := r.db.Create(&user).Error

	return user,err
}
func (r *repository) FindUser() ([]models.User,error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}
func (r *repository) GetUserById(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

