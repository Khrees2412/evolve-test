package database

import (
	"evolve/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	DbUri := os.Getenv("POSTGRES_URI")
	var err error
	DB, err = gorm.Open(postgres.Open(DbUri), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	log.Println("Database connected...")

	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Println(err)
	}
}
