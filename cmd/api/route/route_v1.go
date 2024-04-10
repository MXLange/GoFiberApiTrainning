package route

import (
	"github.com/api/interface/cmd/api/v1/block/user"
	"github.com/gofiber/fiber/v2"
)

func RouteV1(v1 fiber.Router) {
	userRoutes := v1.Group("/users")
	user.UserRoutes(userRoutes)
}
