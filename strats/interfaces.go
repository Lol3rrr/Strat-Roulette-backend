package strats

// SessionInterface exposes all methods that should be exposed by this packet
// and helps with abstracting the implementation away
type SessionInterface interface {
	AddStrat(string, string, Site, []GameMode) error
	GetRandomStrat(Site, GameMode) (Strat, error)
}
