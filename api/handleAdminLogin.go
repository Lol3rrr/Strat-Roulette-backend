package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"
)

func (s *session) handleAdminLogin(ctx *fiber.Ctx) {
	var loginReq loginRequest
	err := ctx.BodyParser(&loginReq)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return
	}

	userSession, err := s.Auth.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		ctx.SendStatus(http.StatusUnauthorized)
		return
	}

	cookieExpiration := time.Unix(userSession.GetExpiration(), 0).Add(1 * time.Hour)

	sessionCookie := new(fiber.Cookie)
	sessionCookie.Name = "sessionID"
	sessionCookie.Value = userSession.GetSessionID()
	sessionCookie.Expires = cookieExpiration
	sessionCookie.Secure = true

	ctx.Cookie(sessionCookie)
	ctx.SendStatus(http.StatusOK)
}
