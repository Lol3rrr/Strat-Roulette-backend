package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp"
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

	sessionCookie := new(fasthttp.Cookie)
	sessionCookie.SetKey("sessionID")
	sessionCookie.SetValue(userSession.GetSessionID())
	sessionCookie.SetExpire(cookieExpiration)
	sessionCookie.SetSecure(true)
	sessionCookie.SetSameSite(fasthttp.CookieSameSiteNoneMode)

	ctx.Fasthttp.Response.Header.SetCookie(sessionCookie)
	ctx.SendStatus(http.StatusOK)
}
