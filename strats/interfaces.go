package strats

// SessionInterface exposes all methods that should be exposed by this packet
// and helps with abstracting the implementation away
type SessionInterface interface {
	// AddStrat takes the given params and inserts a new Strat with these settings into the Database
	AddStrat(Name string, Description string, PlayerSite Site, Modes []GameMode) error

	// GetRandomStrat returns a random Strat from the Database, that matches the given criteria
	GetRandomStrat(PlayerSite Site, Mode GameMode) (Strat, error)

	// GetRandomStrat returns a single Strat with the given ID
	GetStrat(id string) (Strat, error)

	// DeleteStrat removes a single Strat with the given ID
	DeleteStrat(id string) error
}
