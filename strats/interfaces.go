package strats

// SessionInterface exposes all methods that should be exposed by this packet
// and helps with abstracting the implementation away
type SessionInterface interface {
	// AddStrat takes the given params and inserts a new Strat with these settings into the Database
	AddStrat(string, string, Site, []GameMode) error
	// GetRandomStrat returns a random Strat from the Database, that matches the given criteria
	GetRandomStrat(Site, GameMode) (Strat, error)
}
