package auth

// GetRole simply returns the Role of the current UserSession
func (userS *userSession) GetRole() Role {
	return userS.UserRole
}
