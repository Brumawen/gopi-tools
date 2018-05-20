# gopi-tools #

A variety of GPIO utilities for your Raspberry Pi.

**Documentation:** [![GoDoc](https://godoc.org/github.com/brumawen/gopi-tools?status.svg)](https://godoc.org/github.com/brumawen/gopi-tools)

gopi-tools is a Go Library providing a set of tools that control various hardware devices and sensors for a Raspberry Pi.  

Access to the Raspberry Pi's GPIO pins is provided by the awesome [go-rpio](https://github.com/stianeikeland/go-rpio) toolkit.

## Installation and Usage ##

To install run the following on the command prompt:
```Shell
go get "github.com\brumawen\gopi-tools" 
```

Add the following import to the top of your code:
```go
import "github.com/brumawen/gopi-tools"
```

## Tool Module Descriptions ##

### Led ###

This type provides control over a LED that has been [connected](https://thepihut.com/blogs/raspberry-pi-tutorials/27968772-turning-on-an-led-with-your-raspberry-pis-gpio-pins) to a GPIO pin of your Pi.

![alt text](https://github.com/Brumawen/gopi-tools/blob/master/docs/led.png?raw=true "LED schematic")

```go
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
```

### CharDisplay ###

This type provides control for a [LCD Display.](https://learn.adafruit.com/drive-a-16x2-lcd-directly-with-a-raspberry-pi/overview)

![alt text](https://github.com/Brumawen/gopi-tools/blob/master/docs/chardisplay.png?raw=true "CharDisplay schematic")

```go
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
```

### OneWireTemp ###

This type provides control for a 1-Wire Temperature device like a DS18B20 temperature probe.

![alt text](https://github.com/Brumawen/gopi-tools/blob/master/docs/onewire.png?raw=true "OneWire schematic")

```go
func TestCanReadTemp(t *testing.T) {
	o := OneWireTemp{ID: "28-0516a4c75bff"}
	err := o.Init()
	if err != nil {
		t.Error(err)
	}
	defer o.Close()
	v, err := o.ReadTemp()
	if err != nil {
		t.Error(err)
	}
	if v == 999999 {
		t.Error("Temperature could not be read from the file.")
	} else {
		fmt.Println("Temperature", v)
	}
}
```

### Mcp3008 ###

This type provides control for a 8-Channel ADC IC.

![alt text](https://github.com/Brumawen/gopi-tools/blob/master/docs/mcp3008.png?raw=true "MCP3008 schematic")

A call to the Read method returns a float slice containing the 8 channel values read from the IC.

```go
func TestMcp3008CanReadChannels(t *testing.T) {
	m := Mcp3008{}
	if err := m.Init(); err != nil {
		t.Error(err)
	}
	defer m.Close()

	r := m.Read()
	fmt.Println(r)
}
```

### PIN ###
