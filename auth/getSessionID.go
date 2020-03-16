package auth

// GetSessionID simply returns the sessionID of the current UserSession
func (userS *userSession) GetSessionID() string {
	return userS.SessionID
}