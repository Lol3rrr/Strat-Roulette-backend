package api

import (
	"net/http"
	"testing"

	"github.com/go-playground/assert"
)

func TestInitAPI(t *testing.T) {
	routes := []struct {
		Path   string
		Method string
	}{}

	testSession := session{}

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
