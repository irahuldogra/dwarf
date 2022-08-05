package server

import (
	"shortener/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllRedirects(ctx *fiber.Ctx) error {
	dwarves, err := model.GetAllDwarves()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all dwarves links" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dwarves)
}

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/dwarf", getAllRedirects)

	router.Listen(":8080")

}
