package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		type BodyModel struct {
			Message string `json:"message" xml:"message" form:"message"`
		}

		body := new(BodyModel)

		if err := c.BodyParser(body); err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"message": body.Message,
			"status":  "success",
		})
	})

	app.Listen(":8080")
}
