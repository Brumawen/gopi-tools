package gopitools

import (
	"testing"
	"time"
)

func TestPin(t *testing.T) {
	l := Led{GpioLed: 18}
	if err := l.Init(); err != nil {
		t.Error(err)
	}
	defer l.Close()

	l.On()
	time.Sleep(2 * time.Second)
	l.Off()
}

func TestTurnOnPin(t *testing.T) {
	l := Led{GpioLed: 22}
	if err := l.Init(); err != nil {
		t.Error(err)
	}
	defer l.Close()

	l.On()
}

func TestTurnOffPin(t *testing.T) {
	l := Led{GpioLed: 22}
	if err := l.Init(); err != nil {
		t.Error(err)
	}
	defer l.Close()

	l.Off()
}
