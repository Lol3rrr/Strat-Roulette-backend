package api

import "github.com/gofiber/fiber"

func (s *session) init() *fiber.App {
	app := fiber.New()

	return app
}
