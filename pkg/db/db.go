package db

import (
	"log"

	"github.com/Bryan-BC/go-course-microservice/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DataBase *gorm.DB
}

func Init(url string) DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Panicf("Error opening database: %s \n", err)
	}

	db.AutoMigrate(&models.Course{})
	return DB{DataBase: db}
}
