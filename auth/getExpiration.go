package auth

func (userS *userSession) GetExpiration() int64 {
	return userS.Expiration
}
