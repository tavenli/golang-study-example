package main

import (
	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	html "github.com/gofiber/template/html/v2"
)

func main() {

	ViewsEngine := html.New("./views", ".html")

	//app := fiber.New()

	app := fiber.New(fiber.Config{
		Views: ViewsEngine,
	})

	app.Static("/static", "./static")

	app.Use(cors.New())

	app.Get("/login", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("login", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	app.Get("/txt", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/index", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Listen(":3000")
}
