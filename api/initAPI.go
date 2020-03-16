package api

import "github.com/gofiber/fiber"

func (s *session) init() *fiber.App {
	app := fiber.New()

	app.Get("/strat/random", s.handleGetRandomStrat)
	app.Get("/strat/single", s.handleGetSingleStrat)

	app.Post("/admin/login", s.handleAdminLogin)
	app.Use("/admin/strat/", s.middlewareAuth)
	app.Get("/admin/strat/all", s.handleGetAllStrats)
	app.Post("/admin/strat/add", s.handleAddStrat)
	app.Post("/admin/strat/delete", s.handleDeleteStrat)

	return app
}
