package presenter

import "github.com/gofiber/fiber/v2"

func (p *presenter) GetUsers(c *fiber.Ctx) error {
	return c.SendString("user list")
}

func (p *presenter) GetUserDetail(c *fiber.Ctx) error {
	return c.SendString("user detail")
}
