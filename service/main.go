package main

import (
	"dwarf/model"
	"dwarf/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.Setup()
	server.SetupAndListen()
}
