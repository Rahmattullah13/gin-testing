package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dbDsn := "host=localhost user=postgres password=rahmat011099 dbname=gin_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	} else {
		fmt.Println("connected to db")
		DB = db
	}
}
