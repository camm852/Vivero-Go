package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// var DSN = `host={$host}postgres user=root password=root dbname=vivero port=5432`
var DB *gorm.DB

func Connection() {
	host := os.Getenv("DB_HOST")
	userDb := os.Getenv("DB_USER")
	passwordDb := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, userDb, passwordDb, dbName, port)
	fmt.Println(DSN)

	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}

}
