package gopitools

import (
	"fmt"
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

func TestTurnOffOnClose(t *testing.T) {
	l := Led{GpioLed: 18, TurnOffOnClose: true}
	defer l.Close()
	l.On()
	time.Sleep(1 * time.Second)
}

func TestFlashLed(t *testing.T) {
	l := Led{GpioLed: 18}
	if err := l.Init(); err != nil {
		t.Error(err)
	}
	defer l.Close()

	fmt.Println("Flashing every 250 ms")
	l.Flash(250)
	time.Sleep(2 * time.Second)
	fmt.Println("Flashing every 50 ms")
	l.Flash(50)
	time.Sleep(2 * time.Second)
	l.Off()
}
