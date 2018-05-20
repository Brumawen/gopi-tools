package gopitools

// PiModule defines the interface for a component module.
type PiModule interface {
	Close()
	Init() error
}
