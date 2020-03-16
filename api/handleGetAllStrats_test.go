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

func TestHandleGetAllStrats(t *testing.T) {
	tables := []struct {
		Name             string
		InputSession     session
		InputStratsValue []strats.Strat
		InputStratsError error
		ResponseCode     int
		ResponseBody     string
	}{
		{
			Name:         "Valid Input",
			InputSession: session{},
			InputStratsValue: []strats.Strat{
				{
					ID:          "testID1",
					Name:        "testName1",
					Description: "testDescription1",
					PlayerSite:  strats.Attacker,
					Modes: []strats.GameMode{
						strats.Bomb,
					},
				},
			},
			InputStratsError: nil,
			ResponseCode:     http.StatusOK,
			ResponseBody:     `{"strats":[{"id":"testID1","name":"testName1","description":"testDescription1","site":"attacker","modes":["bomb"]}]}`,
		},
		{
			Name:             "Strat returns error",
			InputSession:     session{},
			InputStratsValue: []strats.Strat{},
			InputStratsError: errors.New("testError"),
			ResponseCode:     http.StatusInternalServerError,
			ResponseBody:     "Internal Server Error",
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inStratsValue := table.InputStratsValue
		inStratsError := table.InputStratsError
		responseCode := table.ResponseCode
		responseBody := table.ResponseBody

		inSession.Strats = &mockStrats{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "GetAllStrats",
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
			app.Get("/test", inSession.handleGetAllStrats)

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
