package gopitools

import (
	"os"
	"os/exec"
	"strconv"
)

// Pin allows you to control a GPIO Pin
type Pin struct {
	// Public values
	GpioNo         int
	TurnOffOnClose bool
}

// Close releases and unmaps GPIO memory.
func (l *Pin) Close() {
	if l.TurnOffOnClose {
		l.Off()
	}
}

// Init initializes the Pin ready for use.
func (l *Pin) Init() error {
	_, err := os.Stat("gpiopin.py")
	return err
}

// On turns the Pin on
func (l *Pin) On() error {
	if err := l.Init(); err != nil {
		return err
	}
	return exec.Command("python", "gpiopin.py", "-n", strconv.Itoa(l.GpioNo), "-a", "on").Run()
}

// Off turns the Pin off
func (l *Pin) Off() error {
	if err := l.Init(); err != nil {
		return err
	}
	return exec.Command("python", "gpiopin.py", "-n", strconv.Itoa(l.GpioNo), "-a", "off").Run()
}
