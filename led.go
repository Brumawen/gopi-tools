package gopitools

import (
	"github.com/stianeikeland/go-rpio"
)

// Led allows you to control a GPIO connected LED
type Led struct {
	// Public values
	GpioLed int
	// Private values
	isInitialized bool
	pinLed        rpio.Pin
}

// Close releases and unmaps GPIO memory.
func (l *Led) Close() {
	if l.isInitialized {
		rpio.Close()
		l.isInitialized = false
	}
}

// Init initializes the LED ready for use.
func (l *Led) Init() error {
	if l.isInitialized {
		return nil
	}
	// Set the default GPIO pin number
	if l.GpioLed <= 0 {
		l.GpioLed = 18
	}

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		return err
	}
	l.pinLed = rpio.Pin(l.GpioLed)
	l.pinLed.Output()
	return nil
}

// On turns the LED on
func (l *Led) On() error {
	if !l.isInitialized {
		err := l.Init()
		if err != nil {
			return err
		}
	}
	l.pinLed.High()
	return nil
}

// Off turns the LED off
func (l *Led) Off() error {
	if !l.isInitialized {
		err := l.Init()
		if err != nil {
			return err
		}
	}
	l.pinLed.Low()
	return nil
}

// Toggle turns the LED on, if it's off, or off, if it's on.
func (l *Led) Toggle() error {
	if !l.isInitialized {
		err := l.Init()
		if err != nil {
			return err
		}
	}
	l.pinLed.Toggle()
	return nil
}
