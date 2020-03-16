package auth

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestMockUserSession(t *testing.T) {
	mockSession := MockUserSession{
		Mock: mock.Mock{
			ExpectedCalls: []*mock.Call{
				&mock.Call{
					Method: "GetRole",
					ReturnArguments: mock.Arguments{
						Admin,
					},
				},
				&mock.Call{
					Method: "GetSessionID",
					ReturnArguments: mock.Arguments{
						"testSessionID",
					},
				},
				&mock.Call{
					Method: "GetExpiration",
					ReturnArguments: mock.Arguments{
						int64(123),
					},
				},
			},
		},
	}

	t.Run("GetRole", func(t *testing.T) {
		mockSession.GetRole()
	})
	t.Run("GetSessionID", func(t *testing.T) {
		mockSession.GetSessionID()
	})
	t.Run("GetExpiration", func(t *testing.T) {
		mockSession.GetExpiration()
	})
}
