package main

import (
	"dwarf/model"
	"dwarf/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
