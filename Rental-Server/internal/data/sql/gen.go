package sql

import (
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//go:generate  gentool -c "./gen.tool"

func Db() {
	// Connect to your database

	dsn := "host=localhost user=admin password=admin dbname=main port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	var user model.RentalUser

	db.First(&user)

	log.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)

}
