package api

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (s *session) handleAddStrat(ctx *fiber.Ctx) {
	var addStrat addStratRequest
	err := ctx.BodyParser(&addStrat)
	if err != nil {
		logrus.Errorf("Could not get Add-Strat Body: %v \n", err)
		ctx.SendStatus(http.StatusBadRequest)
		return
	}

	err = s.Strats.AddStrat(
		addStrat.Name,
		addStrat.Description,
		addStrat.Site,
		addStrat.Modes,
	)
	if err != nil {
		logrus.Errorf("Could not Add Strat: %v \n", err)
		ctx.SendStatus(http.StatusInternalServerError)
		return
	}

	ctx.SendStatus(http.StatusOK)
}
