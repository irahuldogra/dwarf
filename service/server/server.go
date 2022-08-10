package server

import (
	"dwarf/model"
	"dwarf/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/goccy/go-json"

	"github.com/gofiber/swagger"
)

func redirect(ctx *fiber.Ctx) error {
	dwarfURL := ctx.Params("redirect")
	dwarf, err := model.FindByDwarfUrl(dwarfURL)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not find dwarf" + err.Error(),
		})
	}

	dwarf.Clicked += 1
	err = model.UpdateDwarf(dwarf)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return ctx.Redirect(dwarf.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllRedirects(ctx *fiber.Ctx) error {
	dwarves, err := model.GetAllDwarves()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all dwarves links" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dwarves)
}

func getDwarf(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id" + err.Error(),
		})
	}

	dwarf, err := model.GetDwarf(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retreive dwarf from db" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dwarf)
}

func createDwarf(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var dwarf model.Dwarf
	err := ctx.BodyParser(&dwarf)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json" + err.Error(),
		})
	}

	if dwarf.Random {
		dwarf.Dwarf = utils.RandomURL(8)
	}

	dwarf, err = model.CreateDwarf(dwarf)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create dwarf in db" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dwarf)
}

func updateDwarf(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var dwarf model.Dwarf

	err := ctx.BodyParser(&dwarf)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not pase json" + err.Error(),
		})
	}

	err = model.UpdateDwarf(dwarf)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update dwarf link in db" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dwarf)
}

func deleteDwarf(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id from url" + err.Error(),
		})
	}

	err = model.DeleteDwarf(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not delete dwarf from db" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "dwarf deleted",
	})
}

func setupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	if os.Getenv("ENABLE_LOGGER") != "" {
		app.Use(logger.New())
	}
}

func SetupAndListen() {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	setupMiddleware(app)

	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	v1.Get("/swagger/*", swagger.HandlerDefault)

	v1.Get("/r/:redirect", redirect)

	v1.Get("/dwarfs", getAllRedirects)
	v1.Get("/dwarfs/:id", getDwarf)
	v1.Post("/dwarfs", createDwarf)
	v1.Patch("/dwarfs", updateDwarf)
	v1.Delete("/dwarfs/:id", deleteDwarf)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))

}
