package auth

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestMockAuthSession(t *testing.T) {
	mockSession := MockAuthSession{
		Mock: mock.Mock{
			ExpectedCalls: []*mock.Call{
				&mock.Call{
					Method: "GetUserSession",
					ReturnArguments: mock.Arguments{
						&userSession{},
						nil,
					},
				},
				&mock.Call{
					Method: "Login",
					ReturnArguments: mock.Arguments{
						&userSession{},
						nil,
					},
				},
			},
		},
	}

	t.Run("GetUserSession", func(t *testing.T) {
		mockSession.GetUserSession("")
	})
	t.Run("Login", func(t *testing.T) {
		mockSession.Login("", "")
	})
}
