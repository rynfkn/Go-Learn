package migrate

import (
	"log"

	"github.com/rynfkn/crud-app/initializers"
	"github.com/rynfkn/crud-app/models"
)

func Migrate() error {
	err := initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully")
	return nil
}