package gopitools

import "testing"
import "time"
import "fmt"

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
