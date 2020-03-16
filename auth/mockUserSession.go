package auth

import "github.com/stretchr/testify/mock"

// MockUserSession is a simply mock implementation
type MockUserSession struct {
	mock.Mock
}

// GetRole is needed to comply with the interface
func (m *MockUserSession) GetRole() Role {
	args := m.Called()
	return args.Get(0).(Role)
}

// GetSessionID is needed to comply with the interface
func (m *MockUserSession) GetSessionID() string {
	args := m.Called()
	return args.String(0)
}

// GetExpiration is needed to comply with the interface
func (m *MockUserSession) GetExpiration() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}
