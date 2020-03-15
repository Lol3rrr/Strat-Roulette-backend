package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strat-roulette-backend/auth"
	"strconv"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleAdminLogin(t *testing.T) {
	tables := []struct {
		Name           string
		InputSession   session
		InputAuthValue auth.UserSessionInterface
		InputAuthError error
		InputBody      map[string]interface{}
		ResponseCode   int
		ResponseCookie bool
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputAuthValue: &auth.MockUserSession{
				Mock: mock.Mock{
					ExpectedCalls: []*mock.Call{
						&mock.Call{
							Method: "GetSessionID",
							ReturnArguments: mock.Arguments{
								"testSessionID",
							},
						},
						&mock.Call{
							Method: "GetExpiration",
							ReturnArguments: mock.Arguments{
								int64(100),
							},
						},
					},
				},
			},
			InputAuthError: nil,
			InputBody: map[string]interface{}{
				"username": "testAdmin",
				"password": "testPassword",
			},
			ResponseCode:   http.StatusOK,
			ResponseCookie: true,
		},
		{
			Name:           "Auth returns error",
			InputSession:   session{},
			InputAuthValue: &auth.MockUserSession{},
			InputAuthError: errors.New("testError"),
			InputBody: map[string]interface{}{
				"username": "testAdmin",
				"password": "testPassword",
			},
			ResponseCode:   http.StatusUnauthorized,
			ResponseCookie: false,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inAuthValue := table.InputAuthValue
		inAuthError := table.InputAuthError
		inBody := table.InputBody
		responseCode := table.ResponseCode
		responseCookie := table.ResponseCookie

		inSession.Auth = &auth.MockAuthSession{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "Login",
						ReturnArguments: mock.Arguments{
							inAuthValue,
							inAuthError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/test", inSession.handleAdminLogin)

			bodyContent, err := json.Marshal(inBody)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(bodyContent))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Content-Length", strconv.Itoa(len(bodyContent)))
			if err != nil {
				t.Fatal(err)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}

			respCode := resp.StatusCode
			assert.Equal(t, responseCode, respCode)

			hasSessionCookie := false
			for _, cookie := range resp.Cookies() {
				if cookie.Name == "sessionID" {
					hasSessionCookie = true
					break
				}
			}

			assert.Equal(t, responseCookie, hasSessionCookie)
		})
	}
}
