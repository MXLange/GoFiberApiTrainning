package main

import (
	"os"

	"github.com/api/interface/cmd/api/route"
	"github.com/api/interface/internal/initializer"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializer.LoadEnv()
}

func main() {
	app := fiber.New(fiber.Config{
		// Prefork: true,
		AppName: "ApiV1",
	})

	route.Routes(app)

	servicePort := os.Getenv("PORT")
	app.Listen(servicePort)
}
