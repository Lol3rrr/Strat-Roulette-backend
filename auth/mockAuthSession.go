package auth

import "github.com/stretchr/testify/mock"

// MockAuthSession is a simply mock implementation
type MockAuthSession struct {
	mock.Mock
}

// GetUserSession is needed to comply with the interface
func (m *MockAuthSession) GetUserSession(sessionID string) (UserSessionInterface, error) {
	args := m.Called()
	return args.Get(0).(UserSessionInterface), args.Error(1)
}

// Login is needed to comply with the interface
func (m *MockAuthSession) Login(username, password string) (UserSessionInterface, error) {
	args := m.Called()
	return args.Get(0).(UserSessionInterface), args.Error(1)
}
