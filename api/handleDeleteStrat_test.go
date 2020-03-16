package api

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleDeleteStrat(t *testing.T) {
	tables := []struct {
		Name             string
		InputSession     session
		InputStratsError error
		ResponseCode     int
	}{
		{
			Name:             "Valid Input",
			InputSession:     session{},
			InputStratsError: nil,
			ResponseCode:     http.StatusOK,
		},
		{
			Name:             "Strats returns error",
			InputSession:     session{},
			InputStratsError: errors.New("testError"),
			ResponseCode:     http.StatusInternalServerError,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inStratsError := table.InputStratsError
		responseCode := table.ResponseCode

		inSession.Strats = &mockStrats{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "DeleteStrat",
						ReturnArguments: mock.Arguments{
							inStratsError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/test", inSession.handleDeleteStrat)

			req, err := http.NewRequest(http.MethodPost, "/test", nil)
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
