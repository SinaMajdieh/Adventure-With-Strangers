package pages

// Static pages handler

import "github.com/gofiber/fiber/v2"

// Let react front end take care of routes
func ReactDefaultHandler(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
