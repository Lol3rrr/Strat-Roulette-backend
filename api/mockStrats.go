package api

import (
	"strat-roulette-backend/strats"

	"github.com/stretchr/testify/mock"
)

type mockStrats struct {
	mock.Mock
}

func (m *mockStrats) AddStrat(string, string, strats.Site, []strats.GameMode) error {
	return m.Called().Error(0)
}

func (m *mockStrats) GetRandomStrat(strats.Site, strats.GameMode) (strats.Strat, error) {
	args := m.Called()
	return args.Get(0).(strats.Strat), args.Error(1)
}

func (m *mockStrats) GetStrat(id string) (strats.Strat, error) {
	args := m.Called()
	return args.Get(0).(strats.Strat), args.Error(1)
}

func (m *mockStrats) DeleteStrat(id string) error {
	return m.Called().Error(0)
}
