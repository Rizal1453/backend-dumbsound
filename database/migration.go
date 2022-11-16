package database

import (
	"dumbsound/models"
	"dumbsound/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
	&models.User{},
	&models.Artis{},
	&models.Song{},
	&models.Transaction{},
	 )
	if err!= nil{
		fmt.Println(err)
		panic("failed to migrate")
	}
fmt.Println("Migration success")
}