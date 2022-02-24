package main

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/interfaces/rest"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rest.Init()
}
