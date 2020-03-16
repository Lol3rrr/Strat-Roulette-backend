package api

import (
	"net/http"
	"strat-roulette-backend/auth"

	"github.com/gofiber/fiber"
)

func (s *session) middlewareAuth(ctx *fiber.Ctx) {
	sessionID := ctx.Cookies("sessionID")
	if len(sessionID) <= 0 {
		ctx.SendStatus(http.StatusUnauthorized)
		return
	}

	userSession, err := s.Auth.GetUserSession(sessionID)
	if err != nil {
		ctx.SendStatus(http.StatusUnauthorized)
		return
	}

	if userSession.GetRole() != auth.Admin {
		ctx.SendStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
