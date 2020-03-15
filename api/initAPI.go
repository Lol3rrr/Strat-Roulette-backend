package api

import "github.com/gofiber/fiber"

func (s *session) init() *fiber.App {
	app := fiber.New()

	app.Get("/strat/random", s.handleGetRandomStrat)
	app.Get("/strat/single", s.handleGetSingleStrat)

	app.Post("/strat/add", s.handleAddStrat)

	return app
}
