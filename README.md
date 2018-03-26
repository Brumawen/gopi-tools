# gopi-tools #

A variety of GPIO utilities for your Raspberry Pi.

**Documentation:** [![GoDoc](https://godoc.org/github.com/brumawen/gopi-tools?status.svg)](https://godoc.org/github.com/brumawen/gopi-tools)

gopi-tools is a Go Library providing a set of tools that control various hardware devices and sensors for a Raspberry Pi.  

Access to the Raspberry Pi's GPIO pins is provided by the awesome [go-rpio](https://github.com/stianeikeland/go-rpio) toolkit.

## Installation and Usage ##

To install run the following on the command prompt:
'''Shell
go get "github.com\brumawen\gopi-tools" 
'''

Add the following import to the top of your code:
'''go
import "github.com/brumawen/gopi-tools"
'''

## Tool Module Descriptions ##

### Led ###

This type provides control over a LED that has been [connected](https://thepihut.com/blogs/raspberry-pi-tutorials/27968772-turning-on-an-led-with-your-raspberry-pis-gpio-pins) to the GPIO pins of your Pi.

'''go
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
'''

### CharDisplay ###

This type provides control for a [LCD Display.](https://learn.adafruit.com/drive-a-16x2-lcd-directly-with-a-raspberry-pi/overview)

'''go
func TestDisplayMessage(t *testing.T) {
	d := CharDisplay{
		GpioRS: 21,
		GpioEN: 20,
		GpioD4: 26,
		GpioD5: 19,
		GpioD6: 13,
		GpioD7: 6,
		Lines:  2,
		Cols:   16,
	}
	err := d.Init()
	if err != nil {
		t.Error(err)
	}
	defer d.Close()
	d.Message("Hello\nWorld")
	time.Sleep(2 * time.Second)
	d.Clear()
}
'''