package main

import (
	"log"

	"web-api/Models/DatabaseModels"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to PostgreSQL database
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&DatabaseModels.User{}, &DatabaseModels.Blog{}, &DatabaseModels.Comment{}, &DatabaseModels.Tag{})
	if err != nil {
		log.Fatal(err)
	}
}
