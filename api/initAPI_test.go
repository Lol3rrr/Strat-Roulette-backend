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
		{
			Path:   "/strat/single",
			Method: http.MethodGet,
		},
		{
			Path:   "/admin/login",
			Method: http.MethodPost,
		},
		{
			Path:   "/admin/strat/all",
			Method: http.MethodGet,
		},
		{
			Path:   "/admin/strat/add",
			Method: http.MethodPost,
		},
		{
			Path:   "/admin/strat/delete",
			Method: http.MethodPost,
		},
	}

	testSession := session{
		Strats: &mockStrats{
			mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "GetRandomStrat",
						ReturnArguments: mock.Arguments{
							strats.Strat{},
							nil,
						},
					},
					&mock.Call{
						Method: "GetStrat",
						ReturnArguments: mock.Arguments{
							strats.Strat{},
							nil,
						},
					},
					&mock.Call{
						Method: "AddStrat",
						ReturnArguments: mock.Arguments{
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
