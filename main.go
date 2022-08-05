package main

import (
	"shortener/model"
	"shortener/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
