package api

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (s *session) handleGetAllStrats(ctx *fiber.Ctx) {
	strats, err := s.Strats.GetAllStrats()
	if err != nil {
		logrus.Errorf("Could not load All-Strats: %v \n", err)
		ctx.SendStatus(http.StatusInternalServerError)
		return
	}

	resp := allStratsResponse{
		Strats: strats,
	}

	err = ctx.JSON(resp)
	if err != nil {
		logrus.Errorf("Could not marshal All-Strats-Response: %v \n", err)
	}
}
