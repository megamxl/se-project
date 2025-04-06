package sql

import (
	"context"
	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/repos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

//go:generate  gentool -c "./gen.tool"

func Db() {
	// Connect to your database

	dsn := "host=localhost user=admin password=admin dbname=main port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	use := query.Use(db)
	ctx := context.Background()

	userRepo := repos.UserRepo{
		Q:   use,
		Ctx: ctx,
	}

	// Temporary Seed
	// Create test user if not already present
	testEmail := "test@test.com"
	existingUser, err := userRepo.GetUserByEmail(testEmail)
	if err == nil {
		slog.Info("✅ User already exists", "user", existingUser)
	} else {
		newUser := dataInt.RentalUser{
			Name:     "test",
			Email:    testEmail,
			Password: "Admin1234!",
		}

		createdUser, err := userRepo.SaveUser(newUser)
		if err != nil {
			log.Fatal("❌ Could not create test user:", err)
		}
		slog.Info("✅ Test user created", "user", createdUser)
	}
}
