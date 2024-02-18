package user

import "github.com/gofiber/fiber/v2"

type UserHandler struct {
	useCase IUserUseCase
}

func NewUserHandler(useCase IUserUseCase) *UserHandler {
	return &UserHandler{useCase}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var input CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	user, err := uh.useCase.CreateUser(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (uh *UserHandler) Authenticate(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var input AuthenticateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	user, err := uh.useCase.Authenticate(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(user)
}
