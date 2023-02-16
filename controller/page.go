package controller

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"name": "Miftah Firdaus",
	})
}

func PageLogin(c *fiber.Ctx) error {
	return c.Render("loginPage", nil)
}

func PageRegister(c *fiber.Ctx) error {
	return c.Render("registrationPage", nil)
}
