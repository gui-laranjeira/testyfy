package guests

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type GuestHandler struct {
	useCase IGuestUseCase
}

func NewGuestHandler(useCase IGuestUseCase) *GuestHandler {
	return &GuestHandler{useCase}
}

func (gh *GuestHandler) CreateGuest(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var input CreateGuestInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	guest, err := gh.useCase.CreateGuest(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusCreated).JSON(guest)
}

func (gh *GuestHandler) GetGuestById(c *fiber.Ctx) error {
	id := c.Params("id")

	guest, err := gh.useCase.FindGuestById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(guest)
}

func (gh *GuestHandler) GetGuestByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")

	guests, err := gh.useCase.FindGuestByUserId(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(guests)
}

func (gh *GuestHandler) GetGuestByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	guest, err := gh.useCase.FindGuestByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(guest)
}

func (gh *GuestHandler) UpdateGuest(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id := c.Params("id")

	var input UpdateGuestInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	input.ID = id

	guest, err := gh.useCase.UpdateGuest(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(guest)
}

func (gh *GuestHandler) DeleteGuest(c *fiber.Ctx) error {
	id := c.Params("id")

	rows, err := gh.useCase.DeleteGuest(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON("rows affected: " + strconv.Itoa(rows))
}
