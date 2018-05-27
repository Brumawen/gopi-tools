package gopitools

import (
	"errors"
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
	return l.runAction("on")
}

// Off turns the Pin off
func (l *Pin) Off() error {
	return l.runAction("off")
}

func (l *Pin) runAction(a string) error {
	if err := l.Init(); err != nil {
		return err
	}
	out, err := exec.Command("python", "gpiopin.py", "-n", strconv.Itoa(l.GpioNo), "-a", a).CombinedOutput()
	if err != nil {
		msg := string(out)
		if msg != "" {
			return errors.New(msg)
		}
	}
	return err
}
