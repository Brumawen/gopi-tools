package gopitools

type PiModule interface {
	Close()
	Init() error
}
