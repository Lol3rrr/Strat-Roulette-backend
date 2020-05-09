package auth

import (
	"github.com/Lol3rrr/mongovault"
)

// Role indicates the Role of a given Session
type Role string

const (
	// Admin simply holds the value for the Admin-Role
	Admin Role = "admin"
)

type session struct {
	Database        mongovault.DB
	AdminUsername   string
	AdminPassword   string
	SessionDuration int64
}

type userSession struct {
	SessionID  string `bson:"sessionID"`
	UserRole   Role   `bson:"userRole"`
	Created    int64  `bson:"created"`
	Expiration int64  `bson:"expiration"`
}
