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

var Database *gorm.DB

func InitDB() {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	var err error
	Database, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connected to database")
	}

	if err := Database.AutoMigrate(&models.Responce{}); err != nil {
		log.Fatal("could not migrate")
	}

}
