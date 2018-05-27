package gopitools

import (
	"os"
	"os/exec"
	"strconv"
)

// Led allows you to control a GPIO connected LED
type Led struct {
	// Public values
	GpioLed        int
	TurnOffOnClose bool

	flashInterval int
}

// Close releases and unmaps GPIO memory.
func (l *Led) Close() {
	if l.TurnOffOnClose {
		l.Off()
	}
}

// Init initializes the LED ready for use.
func (l *Led) Init() error {
	_, err := os.Stat("gpiopin.py")
	return err
}

// On turns the LED on
func (l *Led) On() error {
	if err := l.Init(); err != nil {
		return err
	}
	return exec.Command("python", "gpiopin.py", "-n", strconv.Itoa(l.GpioLed), "-a", "on").Run()
}

// Off turns the LED off
func (l *Led) Off() error {
	if err := l.Init(); err != nil {
		return err
	}
	return exec.Command("python", "gpiopin.py", "-n", strconv.Itoa(l.GpioLed), "-a", "off").Run()
}

// Toggle turns the LED on, if it's off, or off, if it's on.
func (l *Led) Toggle() error {
	if err := l.Init(); err != nil {
		return err
	}
	return exec.Command("python", "gpiopin.py", "-n", strconv.Itoa(l.GpioLed), "-a", "toggle").Run()
}
