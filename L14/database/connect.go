package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // injects ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"not null"`
	Email      string `gorm:"uniqueIndex;not null"`
	Age        int
}

func ConnectDB() (*gorm.DB, error) {
	databaseUri := os.Getenv("DATABASE_URI")
	if databaseUri == "" {
		return nil, fmt.Errorf("DATABASE_URI environment variable not set")
	}
	db, err := gorm.Open(postgres.Open(databaseUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database")
	return db, nil
}
