package database

import (
	"reflect"

	"github.com/stretchr/testify/mock"
)

// MockDatabase can easily be used as a mock implementation
type MockDatabase struct {
	mock.Mock

	connectError error
	getError     error
	getAllError  error
	insertError  error
	updateError  error
	deleteError  error
}

// Get is needed to compy with the interface
func (m *MockDatabase) Get(filter map[string]interface{}, result interface{}) error {
	args := m.Called()

	resultPtrValue := reflect.ValueOf(result)
	resultValue := reflect.Indirect(resultPtrValue)
	resultValue.Set(reflect.ValueOf(args.Get(0)))

	return m.Called().Error(1)
}

// GetAll is needed to compy with the interface
func (m *MockDatabase) GetAll(query map[string]interface{}, results interface{}) error {
	mockResults := m.Called()

	resultsVal := reflect.ValueOf(results)
	sliceVal := resultsVal.Elem()
	sliceVal.Set(reflect.ValueOf(mockResults.Get(0)))

	resultsVal.Elem().Set(sliceVal.Slice(0, mockResults.Int(1)))

	return mockResults.Error(2)
}

// Insert is needed to compy with the interface
func (m *MockDatabase) Insert(interface{}) error {
	return m.Called().Error(0)
}

// Update is needed to compy with the interface
func (m *MockDatabase) Update(map[string]interface{}, map[string]interface{}) error {
	return m.Called().Error(0)
}

// Delete is needed to compy with the interface
func (m *MockDatabase) Delete(map[string]interface{}) error {
	return m.Called().Error(0)
}
