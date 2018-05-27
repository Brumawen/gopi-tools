package gopitools

import (
	"testing"
	"time"
)

func TestLed(t *testing.T) {
	l := Led{GpioLed: 18}
	if err := l.Init(); err != nil {
		t.Error(err)
	}
	defer l.Close()

	l.On()
	time.Sleep(2 * time.Second)
	l.Off()
}

func TestTurnLedOffOnClose(t *testing.T) {
	l := Led{GpioLed: 18, TurnOffOnClose: true}
	defer l.Close()
	l.On()
	time.Sleep(1 * time.Second)
}
