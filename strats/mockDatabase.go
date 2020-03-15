package strats

import (
	"reflect"

	"github.com/stretchr/testify/mock"
)

type mockDatabase struct {
	mock.Mock

	connectError error
	getError     error
	getAllError  error
	insertError  error
	updateError  error
	deleteError  error
}

func (m *mockDatabase) Get(filter map[string]interface{}, result interface{}) error {
	args := m.Called()

	resPtr := result.(*Strat)
	*resPtr = args.Get(0).(Strat)

	return m.Called().Error(1)
}

func (m *mockDatabase) GetAll(query map[string]interface{}, results interface{}) error {
	mockResults := m.Called()

	resultsVal := reflect.ValueOf(results)
	sliceVal := resultsVal.Elem()
	sliceVal.Set(reflect.ValueOf(mockResults.Get(0)))

	resultsVal.Elem().Set(sliceVal.Slice(0, mockResults.Int(1)))

	return mockResults.Error(2)
}

func (m *mockDatabase) Insert(interface{}) error {
	return m.Called().Error(0)
}

func (m *mockDatabase) Update(map[string]interface{}, map[string]interface{}) error {
	return m.Called().Error(0)
}

func (m *mockDatabase) Delete(map[string]interface{}) error {
	return m.Called().Error(0)
}
