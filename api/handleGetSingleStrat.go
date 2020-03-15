package api

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (s *session) handleGetSingleStrat(ctx *fiber.Ctx) {
	id := ctx.Query("id")

	result, err := s.Strats.GetStrat(id)
	if err != nil {
		logrus.Errorf("Could not load Single-Strat: %v \n", err)
		ctx.SendStatus(http.StatusInternalServerError)
		return
	}

	err = ctx.JSON(result)
	if err != nil {
		logrus.Errorf("Could not marshal Single-Strat-Response: %v \n", err)
	}
}
