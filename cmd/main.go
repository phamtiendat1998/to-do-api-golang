package main

import (
	"log"
	"os"

	"to-do-api-golang/controllers"
	"to-do-api-golang/seed"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func main() {
	var err error
	err = godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error getting env: %v", err)
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")
}
