package strats

import "strat-roulette-backend/database"

// GameMode is a simple enum that stores a Gamemode of a Strat
type GameMode string

const (
	// Bomb is the const for the "Bomb" Gamemode
	Bomb GameMode = "bomb"
	// Hostage is the const for the "Hostage" Gamemode
	Hostage GameMode = "hostage"
	// SecureArea is the const for the "Secure Area" Gamemode
	SecureArea GameMode = "secureArea"
)

// Site is a simple enum that stores the Site of the Strat
type Site string

const (
	// Attacker is the const for the "Attacker" Site
	Attacker Site = "attacker"
	// Defender is the const for the "Defender" Site
	Defender Site = "defender"
)

// Strat describes a single Strategy
type Strat struct {
	ID          string     `json:"id" bson:"id"`
	Name        string     `json:"name" bson:"name"`
	Description string     `json:"description" bson:"description"`
	PlayerSite  Site       `json:"-" bson:"playerSite"`
	Modes       []GameMode `json:"-" bson:"modes"`
}

type session struct {
	Database database.SessionInterface
}