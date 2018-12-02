package process

// Killer is a process killer
type Killer interface {
	Kill() error
}
