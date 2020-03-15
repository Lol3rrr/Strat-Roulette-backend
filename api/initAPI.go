package api

import "github.com/gofiber/fiber"

func (s *session) init() *fiber.App {
	app := fiber.New()

	app.Get("/strat/random", s.handleGetRandomStrat)

	return app
}
