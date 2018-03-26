package gopitools

import "testing"
import "time"

func TestLed(t *testing.T) {
	l := Led{GpioLed: 18}
	err := l.Init()
	if err != nil {
		t.Error(err)
	}
	defer l.Close()

	l.On()
	time.Sleep(2 * time.Second)
	l.Off()
}
