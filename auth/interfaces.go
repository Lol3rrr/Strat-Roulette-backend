package auth

// SessionInterface represents all the functions exposed by this package
type SessionInterface interface {
	// GetUserSession loads the Session associated with the given SessionID
	// returns an error if no session is found
	GetUserSession(sessionID string) (UserSessionInterface, error)

	// Login takes a username and a password and tries to log the user in
	// with the credentials and returns a valid SessionInterface if it worked
	// otherwise returns an error
	Login(username, password string) (UserSessionInterface, error)
}

// UserSessionInterface represents a single Session for a User
type UserSessionInterface interface {
	// GetRole returns the Role of the current Session
	GetRole() Role

	// GetSessionID returns the SessionID associated with the
	// current UserSession
	GetSessionID() string
}
