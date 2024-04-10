package route

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	RouteV1(v1)
}
