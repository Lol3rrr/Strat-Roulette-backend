package strats

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddStrat(t *testing.T) {

	tables := []struct {
		Name             string
		InputSession     session
		InputDBError     error
		InputName        string
		InputDescription string
		InputPlayerSite  Site
		InputModes       []GameMode
		ResultError      bool
	}{
		{
			Name:             "Valid Input",
			InputSession:     session{},
			InputDBError:     nil,
			InputName:        "testName",
			InputDescription: "testDescription",
			InputPlayerSite:  Attacker,
			InputModes: []GameMode{
				Bomb,
			},
			ResultError: false,
		},
		{
			Name:             "DB returns error",
			InputSession:     session{},
			InputDBError:     errors.New("testError"),
			InputName:        "testName",
			InputDescription: "testDescription",
			InputPlayerSite:  Attacker,
			InputModes: []GameMode{
				Bomb,
			},
			ResultError: true,
		},
		{
			Name:             "Empty Name",
			InputSession:     session{},
			InputDBError:     nil,
			InputName:        "",
			InputDescription: "testDescription",
			InputPlayerSite:  Attacker,
			InputModes: []GameMode{
				Bomb,
			},
			ResultError: true,
		},
		{
			Name:             "Empty Description",
			InputSession:     session{},
			InputDBError:     nil,
			InputName:        "testName",
			InputDescription: "",
			InputPlayerSite:  Attacker,
			InputModes: []GameMode{
				Bomb,
			},
			ResultError: true,
		},
		{
			Name:             "Empty Site",
			InputSession:     session{},
			InputDBError:     nil,
			InputName:        "testName",
			InputDescription: "testDescription",
			InputPlayerSite:  "",
			InputModes: []GameMode{
				Bomb,
			},
			ResultError: true,
		},
		{
			Name:             "No Gamemode",
			InputSession:     session{},
			InputDBError:     nil,
			InputName:        "testName",
			InputDescription: "testDescription",
			InputPlayerSite:  Attacker,
			InputModes:       []GameMode{},
			ResultError:      true,
		},
		{
			Name:             "Unknown Site",
			InputSession:     session{},
			InputDBError:     nil,
			InputName:        "testName",
			InputDescription: "testDescription",
			InputPlayerSite:  "testSite",
			InputModes: []GameMode{
				Bomb,
			},
			ResultError: true,
		},
		{
			Name:             "Unknown Mode",
			InputSession:     session{},
			InputDBError:     nil,
			InputName:        "testName",
			InputDescription: "testDescription",
			InputPlayerSite:  Attacker,
			InputModes: []GameMode{
				"testMode",
			},
			ResultError: true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inDBError := table.InputDBError
		inName := table.InputName
		inDescription := table.InputDescription
		inPlayerSite := table.InputPlayerSite
		inModes := table.InputModes
		resultError := table.ResultError

		inSession.Database = &mockDatabase{
			Mock: mock.Mock{
				ExpectedCalls: []*mock.Call{
					&mock.Call{
						Method: "Insert",
						ReturnArguments: mock.Arguments{
							inDBError,
						},
					},
				},
			},
		}

		t.Run(table.Name, func(t *testing.T) {
			outputError := inSession.AddStrat(inName, inDescription, inPlayerSite, inModes)

			if resultError {
				assert.NotNil(t, outputError)
			} else {
				assert.Nil(t, outputError)
			}
		})
	}
}
