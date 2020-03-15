package api

import (
	"net/http"
	"strat-roulette-backend/strats"
	"testing"

	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitAPI(t *testing.T) {
	routes := []struct {
		Path   string
		Method string
	}{
		{
			Path:   "/strat/random",
			Method: http.MethodGet,
		},
	}

	testSession := session{
		&mockStrats{
			mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "GetRandomStrat",
						ReturnArguments: mock.Arguments{
							strats.Strat{},
							nil,
						},
					},
				},
			},
		},
	}

	app := testSession.init()

	for _, route := range routes {
		req, err := http.NewRequest(route.Method, route.Path, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		resStatus := resp.StatusCode
		assert.NotEqual(t, 404, resStatus)
	}
}
