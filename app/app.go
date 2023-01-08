package app

import (
	"fmt"
	"keycloak_api_go/app/database"
	"keycloak_api_go/app/route"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AppStart() {
	loadEnv()
	loadDatabaseConnection()

	runAppServe()
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
	// database.Database.AutoMigrate(&model.Entry{})
}

func runAppServe() {
	router := gin.Default()

	route.AuthRoutes(router)
	route.BlogRoutes(router)

	PORT := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	router.Run(PORT)
	fmt.Println("Server Running on port " + PORT)
}
