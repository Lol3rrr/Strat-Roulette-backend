package api

import (
	"errors"
	"net/http"
	"strat-roulette-backend/auth"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMiddlewareAuth(t *testing.T) {
	tables := []struct {
		Name             string
		InputSession     session
		InputAuthError   error
		InputAuthRole    auth.Role
		InputCookieName  string
		InputCookieValue string
		ResponseCode     int
	}{
		{
			Name:             "Valid Input",
			InputSession:     session{},
			InputAuthError:   nil,
			InputAuthRole:    auth.Admin,
			InputCookieName:  "sessionID",
			InputCookieValue: "testSessionID",
			ResponseCode:     http.StatusOK,
		},
		{
			Name:             "Auth returns error",
			InputSession:     session{},
			InputAuthError:   errors.New("testError"),
			InputAuthRole:    auth.Admin,
			InputCookieName:  "sessionID",
			InputCookieValue: "testSessionID",
			ResponseCode:     http.StatusUnauthorized,
		},
		{
			Name:             "Cookie has wrong name",
			InputSession:     session{},
			InputAuthError:   nil,
			InputAuthRole:    auth.Admin,
			InputCookieName:  "session",
			InputCookieValue: "testSessionID",
			ResponseCode:     http.StatusUnauthorized,
		},
		{
			Name:             "Cookie has empty Value",
			InputSession:     session{},
			InputAuthError:   nil,
			InputAuthRole:    auth.Admin,
			InputCookieName:  "sessionID",
			InputCookieValue: "",
			ResponseCode:     http.StatusUnauthorized,
		},
		{
			Name:             "Session-Role is not Admin",
			InputSession:     session{},
			InputAuthError:   nil,
			InputAuthRole:    auth.Role("testRole"),
			InputCookieName:  "sessionID",
			InputCookieValue: "testSessionID",
			ResponseCode:     http.StatusUnauthorized,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inAuthError := table.InputAuthError
		inAuthRole := table.InputAuthRole
		inCookieName := table.InputCookieName
		inCookieValue := table.InputCookieValue
		responseCode := table.ResponseCode

		inSession.Auth = &auth.MockAuthSession{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "GetUserSession",
						ReturnArguments: mock.Arguments{
							&auth.MockUserSession{
								Mock: mock.Mock{
									ExpectedCalls: []*mock.Call{
										&mock.Call{
											Method: "GetRole",
											ReturnArguments: mock.Arguments{
												inAuthRole,
											},
										},
									},
								},
							},
							inAuthError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			app := fiber.New()
			app.Use(inSession.middlewareAuth)
			app.Get("/test", func(c *fiber.Ctx) {
				c.SendStatus(http.StatusOK)
			})

			req, err := http.NewRequest(http.MethodGet, "/test", nil)
			req.AddCookie(&http.Cookie{
				Name:  inCookieName,
				Value: inCookieValue,
			})
			if err != nil {
				t.Fatal(err)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}

			respCode := resp.StatusCode
			assert.Equal(t, responseCode, respCode)
		})
	}
}
