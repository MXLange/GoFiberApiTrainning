package user

import (
	"fmt"

	db "github.com/api/interface/internal/database"
	"github.com/gofiber/fiber/v2"
)

func (u *User) CreateUserService() *fiber.Error {

	err := db.Connect()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error connecting to database")
	}

	err = u.GetUserByEmailRepo()

	if err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Email already exists")
	}

	err = u.CreateUserRepo()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error creating user")
	}

	fmt.Println("User created", u)

	return nil
}
