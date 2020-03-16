package api

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (s *session) handleDeleteStrat(ctx *fiber.Ctx) {
	id := ctx.Query("id")

	err := s.Strats.DeleteStrat(id)
	if err != nil {
		logrus.Errorf("Could not Delete-Strat: %v \n", err)
		ctx.SendStatus(http.StatusInternalServerError)
		return
	}

	ctx.SendStatus(http.StatusOK)
}
