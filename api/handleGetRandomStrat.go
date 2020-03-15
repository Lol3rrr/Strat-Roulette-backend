package api

import (
	"net/http"
	"strat-roulette-backend/strats"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (s *session) handleGetRandomStrat(ctx *fiber.Ctx) {
	rawSite := ctx.Query("site")
	rawMode := ctx.Query("mode")

	site := strats.Site(rawSite)
	mode := strats.GameMode(rawMode)

	strat, err := s.Strats.GetRandomStrat(site, mode)
	if err != nil {
		logrus.Errorf("Could not load Random-Strat: %v \n", err)
		ctx.SendStatus(http.StatusInternalServerError)
		return
	}

	err = ctx.JSON(strat)
	if err != nil {
		logrus.Errorf("Could not marshal Random-Strat-Response: %v \n", err)
	}
}
