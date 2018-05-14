package gopitools

import (
	"github.com/stianeikeland/go-rpio"
)

// Pin allows you to control a GPIO Pin
type Pin struct {
	// Public values
	GpioNo         int
	TurnOffOnClose bool
	// Private values
	isInitialized bool
	pin           rpio.Pin
}

// Close releases and unmaps GPIO memory.
func (l *Pin) Close() {
	if l.isInitialized {
		if l.TurnOffOnClose {
			l.Off()
		}
		rpio.Close()
		l.isInitialized = false
	}
}

// Init initializes the Pin ready for use.
func (l *Pin) Init() error {
	if l.isInitialized {
		return nil
	}
	// Set the default GPIO pin number
	if l.GpioNo <= 0 {
		l.GpioNo = 18
	}

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		return err
	}
	l.pin = rpio.Pin(l.GpioNo)
	l.pin.Output()
	l.isInitialized = true
	return nil
}

// On turns the Pin on
func (l *Pin) On() error {
	if !l.isInitialized {
		err := l.Init()
		if err != nil {
			return err
		}
	}
	l.pin.High()
	return nil
}

// Off turns the Pin off
func (l *Pin) Off() error {
	if !l.isInitialized {
		err := l.Init()
		if err != nil {
			return err
		}
	}
	l.pin.Low()
	return nil
}
