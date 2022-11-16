package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction)(models.Transaction,error)
	FindTransaction()([]models.Transaction,error)
	GetTransactionById(ID int)(models.Transaction,error)
	UpdateTransaction(status string, ID string) error
	GetByLogin(UserID int) (models.Transaction, error)

}
 
func RepositoryTransaction(db *gorm.DB) *repository{
	return &repository{db}
}
func (r *repository)CreateTransaction(transaction models.Transaction)(models.Transaction,error){
	err := r.db.Create(&transaction).Error
	return transaction,err
}
func (r *repository)FindTransaction()([]models.Transaction,error){
var transaction []models.Transaction
err:= r.db.Preload("User").Find(&transaction).Error

return transaction,err
}
func (r *repository)GetTransactionById(ID int )(models.Transaction,error){
var transaction models.Transaction
err := r.db.First(&transaction,"id=?",ID).Error
return transaction,err
}
func (r *repository) GetByLogin(UserID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Where("user_id = ? AND status_user = ?", UserID, "Active").First(&transaction).Error
	return transaction, err
}
func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		transaction.Status = status
		transaction.Limit = 30
		transaction.StatusUser = "Active"
	}
	if status != transaction.Status && status == "failed" {
		transaction.Status = status
		transaction.Limit = 0
		transaction.StatusUser = "Not Active"
	}

	err := r.db.Save(&transaction).Error

	return err
}