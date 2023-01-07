package app

import (
	"fmt"
	"keycloak_api_go/app/controller"
	"keycloak_api_go/app/database"
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

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	// protectedRoutes := router.Group("/api")
	// protectedRoutes.Use(middleware.MiddlewareAuthKeycloak())
	// protectedRoutes.POST("/entry", controller.AddEntry)
	// protectedRoutes.GET("/entry", controller.GetAllEntries)

	PORT := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	router.Run(PORT)
	fmt.Println("Server Running on port " + PORT)
}
