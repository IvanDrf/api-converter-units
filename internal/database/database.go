package database

import (
	"fmt"
	"log"

	"github.com/IvanDrf/units/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "yourpassword"
	dbname   = "postgres"
	port     = "5432"
	sslmode  = "disable"
)

func InitDB() *gorm.DB {
	database := new(gorm.DB)
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	var err error
	database, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connected to database")
	}

	if err := database.AutoMigrate(&models.Responce{}); err != nil {
		log.Fatal("could not migrate")
	}

	return database

}
