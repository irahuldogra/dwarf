package server

import (
	"dwarf/model"
	"dwarf/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	err = model.CreateDwarf(dwarf)
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

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/dwarf", getAllRedirects)
	router.Get("/dwarf/:id", getDwarf)
	router.Post("/dwarf", createDwarf)
	router.Patch("/dwarf", updateDwarf)
	router.Delete("/dwarf/:id", deleteDwarf)

	router.Listen(":8080")

}
