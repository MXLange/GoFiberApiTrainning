package user

import "github.com/gofiber/fiber/v2"

func UserRoutes(userRoutes fiber.Router) {
	userRoutes.Post("", Create)
	userRoutes.Get("", GetAll)
	userRoutes.Get("/:id", GetByID)
	userRoutes.Put("/:id", Update)
	userRoutes.Delete("/:id", Delete)
}
