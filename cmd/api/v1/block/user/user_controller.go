package user

import "github.com/gofiber/fiber/v2"

func Create(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": []string{"Error parsing body."},
		})
	}

	errArray, err := user.ValidateNewUser()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errArray,
		})
	}

	fiberErr := user.CreateUserService()

	if fiberErr != nil {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"errors": []string{fiberErr.Error()},
		})
	}

	return nil
}
func GetAll(c *fiber.Ctx) error {
	return nil
}
func GetByID(c *fiber.Ctx) error {
	return nil

}
func Update(c *fiber.Ctx) error {
	return nil

}
func Delete(c *fiber.Ctx) error {
	return nil

}
