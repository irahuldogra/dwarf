package dwarfs

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("/r/:redirect", redirect)

	route.Get("/dwarfs", getAllRedirects)
	route.Get("/dwarfs/:id", getDwarf)
	route.Post("/dwarfs", createDwarf)
	route.Patch("/dwarfs", updateDwarf)
	route.Delete("/dwarfs/:id", deleteDwarf)
}
