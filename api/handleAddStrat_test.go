package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleAddStrat(t *testing.T) {
	tables := []struct {
		Name             string
		InputSession     session
		InputStratsError error
		InputBody        map[string]interface{}
		InputSendBody    bool
		ResponseCode     int
	}{
		{
			Name:             "Valid Input",
			InputSession:     session{},
			InputStratsError: nil,
			InputBody:        map[string]interface{}{},
			InputSendBody:    true,
			ResponseCode:     http.StatusOK,
		},
		{
			Name:             "Strats returns error",
			InputSession:     session{},
			InputStratsError: errors.New("testError"),
			InputBody:        map[string]interface{}{},
			InputSendBody:    true,
			ResponseCode:     http.StatusInternalServerError,
		},
		{
			Name:             "No Body",
			InputSession:     session{},
			InputStratsError: nil,
			InputBody:        map[string]interface{}{},
			InputSendBody:    false,
			ResponseCode:     http.StatusBadRequest,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inStratsError := table.InputStratsError
		inBody := table.InputBody
		inSendBody := table.InputSendBody
		responseCode := table.ResponseCode

		inSession.Strats = &mockStrats{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "AddStrat",
						ReturnArguments: mock.Arguments{
							inStratsError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/test", inSession.handleAddStrat)

			bodyContent, err := json.Marshal(inBody)
			if err != nil {
				t.Fatal(err)
			}

			if !inSendBody {
				bodyContent = []byte("")
			}

			reader := bytes.NewReader(bodyContent)

			req, err := http.NewRequest(http.MethodPost, "/test", reader)
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
		})
	}
}
