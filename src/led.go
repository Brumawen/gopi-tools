package gopitools

import (
	"errors"
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
	return l.runAction("on")
}

// Off turns the LED off
func (l *Led) Off() error {
	return l.runAction("off")
}

// Toggle turns the LED on, if it's off, or off, if it's on.
func (l *Led) Toggle() error {
	return l.runAction("toggle")
}

func (l *Led) runAction(a string) error {
	if err := l.Init(); err != nil {
		return err
	}
	out, err := exec.Command("python", "gpiopin.py", "-n", strconv.Itoa(l.GpioLed), "-a", a).CombinedOutput()
	if err != nil {
		msg := string(out)
		if msg != "" {
			return errors.New(msg)
		}
	}
	return err
}
