package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/julhan07/fiber-service/presenter"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Static("/assets", "./public")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	presenter.NewPresenterIndex(app)
	app.Listen(":3000")
}
