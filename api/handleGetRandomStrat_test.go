package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strat-roulette-backend/strats"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleGetRandomStrat(t *testing.T) {
	tables := []struct {
		Name             string
		InputSession     session
		InputStratsValue strats.Strat
		InputStratsError error
		ResponseCode     int
		ResponesBody     string
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputStratsValue: strats.Strat{
				ID:          "testID",
				Name:        "testName",
				Description: "testDescription",
				PlayerSite:  strats.Attacker,
				Modes: []strats.GameMode{
					strats.Bomb,
				},
			},
			InputStratsError: nil,
			ResponseCode:     200,
			ResponesBody:     `{"id":"testID","name":"testName","description":"testDescription","site":"attacker","modes":["bomb"]}`,
		},
		{
			Name:             "Strats returns error",
			InputSession:     session{},
			InputStratsValue: strats.Strat{},
			InputStratsError: errors.New("testError"),
			ResponseCode:     500,
			ResponesBody:     "Internal Server Error",
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inStratsValue := table.InputStratsValue
		inStratsError := table.InputStratsError
		responseCode := table.ResponseCode
		responseBody := table.ResponesBody

		inSession.Strats = &mockStrats{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "GetRandomStrat",
						ReturnArguments: mock.Arguments{
							inStratsValue,
							inStratsError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/test", inSession.handleGetRandomStrat)

			req, err := http.NewRequest(http.MethodGet, "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}

			respCode := resp.StatusCode
			assert.Equal(t, responseCode, respCode)

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, responseBody, string(body))
		})
	}
}
