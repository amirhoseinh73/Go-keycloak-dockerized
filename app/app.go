package app

import (
	"fmt"
	"keycloak_api_go/app/database"
	"log"

	"github.com/joho/godotenv"
)

func AppStart() {
	loadEnv()
	loadDatabaseConnection()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading env file")
	}

	fmt.Println("env file loaded successfully.")
}

func loadDatabaseConnection() {
	database.DBConnect()
}
