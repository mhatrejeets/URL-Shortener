package main

import "github.com/gofiber/fiber/v2"

func main() {
	initDB()
	initRedis()

	app := fiber.New()
	app.Static("/", "./public")

	app.Post("/shorten", createShortURL)

	app.Get("/:shortcode", RedirectToLongURL)

	app.Listen(":3000")
	defer DB.Close()

}
