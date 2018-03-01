package promise

// State type
type State int

// List of available promise states
const (
	Pending State = iota
	Resolved
	Rejected
)
