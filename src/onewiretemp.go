package gopitools

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// OneWireTemp allows you to read the temperature from a One-Wire
// temperature device like a DS188B20 temperature probe.
type OneWireTemp struct {
	// Public values
	ID string
	// Private values
	filePath      string
	isInitialized bool
	re            *regexp.Regexp
}

// Close releases any resources.
func (t *OneWireTemp) Close() {
	if t.isInitialized {
		t.filePath = ""
		t.isInitialized = false
	}
}

// GetDeviceList returns a list of attached devices.
// The first characters before the dash are the family code.
// The characters after the dash are the unique serial number.
// Family codes are documented here:
// https://github.com/owfs/owfs-doc/wiki/1Wire-Device-List
func GetDeviceList() ([]string, error) {
	files, err := filepath.Glob("/sys/bus/w1/devices/*")
	if err != nil {
		return nil, err
	}
	l := []string{}
	for _, p := range files {
		_, file := filepath.Split(p)
		if file != "w1_bus_master1" {
			l = append(l, file)
		}
	}
	return l, err
}

// ReadTemp returns the current temperature in celcius.
func (t *OneWireTemp) ReadTemp() (float64, error) {
	if !t.isInitialized {
		err := t.Init()
		if err != nil {
			return 0, err
		}
	}
	data, err := ioutil.ReadFile(t.filePath)
	if err != nil {
		return 0, err
	}
	m := t.re.FindStringSubmatch(string(data))
	if len(m) >= 2 {
		v, err := strconv.ParseFloat(m[1], 64)
		return v / 1000, err
	}
	return 999999, nil
}

// ReadTempInFahrenheit returns the current temperature in fahrenheit.
func (t *OneWireTemp) ReadTempInFahrenheit() (float64, error) {
	v, err := t.ReadTemp()
	if err == nil {
		v = (v * 9 / 5) + 32
	}
	return v, err
}

// Init initializes the LED ready for use.
func (t *OneWireTemp) Init() error {
	// Check to see if this device is installed and working
	f, err := t.getFilePath()
	if err == nil {
		t.filePath = f
		t.isInitialized = true
	}
	t.re = regexp.MustCompile("t=(\\d{1,})")
	return err
}

func (t *OneWireTemp) getFilePath() (string, error) {
	fileName := "/sys/bus/w1/devices/" + t.ID + "/w1_slave"
	_, err := os.Stat(fileName)
	return fileName, err
}
