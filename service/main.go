package main

import (
	"dwarf/api"
	"dwarf/database"
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

	app := server.Create()

	database.DB.AutoMigrate(&model.Dwarf{})

	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
