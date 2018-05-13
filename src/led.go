package gopitools

import (
	"errors"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// Led allows you to control a GPIO connected LED
type Led struct {
	// Public values
	GpioLed        int
	TurnOffOnClose bool
	// Private values
	isInitialized bool
	pinLed        rpio.Pin
	flashInterval int
}

// Close releases and unmaps GPIO memory.
func (l *Led) Close() {
	if l.isInitialized {
		if l.TurnOffOnClose {
			l.Off()
		}
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
	l.isInitialized = true
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
	l.flashInterval = 0
	l.pinLed.High()
	return nil
}

// Flash turns on the LED for the specified interval (in ms), then turns it off for the same amount of time.
// This is repeated until the Off function is called.
func (l *Led) Flash(interval int) error {
	if interval <= 0 {
		return errors.New("Interval must be greater than 0.")
	}
	if !l.isInitialized {
		err := l.Init()
		if err != nil {
			return err
		}
	}
	if l.flashInterval > 0 {
		// Flash already enabled
		if l.flashInterval != interval {
			l.flashInterval = interval
		}
	} else {
		l.flashInterval = interval
		go l.doFlash()
	}
	return nil
}

func (l *Led) doFlash() {
	for {
		if l.flashInterval == 0 {
			break
		}
		l.pinLed.High()
		if l.flashInterval == 0 {
			break
		}
		time.Sleep(time.Duration(l.flashInterval) * time.Millisecond)
		if l.flashInterval == 0 {
			break
		}
		l.pinLed.Low()
		if l.flashInterval == 0 {
			break
		}
		time.Sleep(time.Duration(l.flashInterval) * time.Millisecond)
	}
}

// Off turns the LED off
func (l *Led) Off() error {
	if !l.isInitialized {
		err := l.Init()
		if err != nil {
			return err
		}
	}
	l.flashInterval = 0
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
