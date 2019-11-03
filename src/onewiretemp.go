package gopitools

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

// OneWireDevice holds the details about a connected device.
type OneWireDevice struct {
	ID       string
	Code     string
	Name     string
	SerialNo string
}

// Close releases any resources.
func (t *OneWireTemp) Close() {
	if t.isInitialized {
		t.filePath = ""
		t.isInitialized = false
	}
}

// Contains returns whether the list of devices contains the specified ID
func Contains(l []OneWireDevice, id string) bool {
	for _, a := range l {
		if a.ID == id {
			return true
		}
	}
	return false
}

// GetDeviceList returns a list of attached devices.
// The first characters before the dash are the family code.
// The characters after the dash are the unique serial number.
// Family codes are documented here:
// https://github.com/owfs/owfs-doc/wiki/1Wire-Device-List
func GetDeviceList() ([]OneWireDevice, error) {
	files, err := filepath.Glob("/sys/bus/w1/devices/*")
	if err != nil {
		return nil, err
	}
	l := []OneWireDevice{}
	for _, p := range files {
		_, file := filepath.Split(p)
		if file != "w1_bus_master1" {
			p := strings.Index(file, "-")
			if p >= 0 {
				fam := file[0:p]
				ser := file[p+1:]
				if fam != "00" {
					l = append(l, OneWireDevice{
						ID:       file,
						Code:     fam,
						Name:     getFamilyName(fam),
						SerialNo: ser,
					})
				}
			}
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

// Init initializes the device ready for use.
func (t *OneWireTemp) Init() error {
	// Check to see if this device is installed and working
	f, err := t.getFilePath()
	if err == nil {
		t.filePath = f
		t.isInitialized = true
	}
	t.re = regexp.MustCompile("t=(\\d{1,})")
	t.isInitialized = true
	return err
}

func (t *OneWireTemp) getFilePath() (string, error) {
	fileName := "/sys/bus/w1/devices/" + t.ID + "/w1_slave"
	_, err := os.Stat(fileName)
	return fileName, err
}

func getFamilyName(fam string) string {
	switch fam {
	case "02":
		return "DS1991 - 1152b Secure Memory"
	case "04":
		return "DS2404 - EconoRAM Time Chip"
	case "05":
		return "DS2405 - Addressable Switch"
	case "06":
		return "DS1993 - 4Kb Memory Button"
	case "08":
		return "DS1992 - 1Kb Memory Button"
	case "0A":
		return "DS1995 - 16Kb Memory Button"
	case "0B":
		return "DS1985,DS2505 - 16Kb Add-only Memory"
	case "0C":
		return "DS1996, 64Kb Memory Button"
	case "0F":
		return "DS1986,DS2506 - 65Kn Add-only Memory"
	case "10":
		return "DS18S20 - Temperature Sensor"
	case "12":
		return "DS2406,DS2407 - Dual Addressable Switch"
	case "14":
		return "DS1971,DS2430A - 256b EEPROM"
	case "16":
		return "DS1954,DS1957 - Coprocessor iButton"
	case "18":
		return "DS1962,DS1963S - 4Kb Monetary Device with SHA"
	case "1A":
		return "DS1963L - 4Kb Monetary Device"
	case "1B":
		return "DS2436 - Battery ID/Monitor"
	case "1C":
		return "DS28E04-100 - 4Kb EEPROM with PIO"
	case "1D":
		return "DS2423 - 4Kb 1Wire RAM with Counter"
	case "1E":
		return "DS2437 - Smart Battery Monitor IC"
	case "1F":
		return "DS2409 - MicroLan Coupler"
	case "20":
		return "DS2450 - Quad ADC"
	case "21":
		return "DS1921G,DS1921H,DS1921Z - Thermochron Loggers"
	case "22":
		return "DS1822 - Econo Digital Thermomenter"
	case "23":
		return "DS1973,DS2433 - 4Kn EEPROM"
	case "24":
		return "SD2415 - Time Chip"
	case "26":
		return "DS2438 - Smart Battery Monitor"
	case "27":
		return "DS2417 - Time Chip"
	case "28":
		return "DS18B20 - Temperature Sensor"
	case "29":
		return "DS2408 - 8-Channel Switch"
	case "2C":
		return "DS2890 - Digital Potentiometer"
	case "2D":
		return "DS1972,DS2431 - 1024Kb Memory"
	case "2E":
		return "DS2770 - Battery Monitor/Charge Controller"
	case "30":
		return "DS2760 - Precision LI+ Batter Monitor"
	case "31":
		return "DS2720 - Single Cell LI+ Protection IC"
	case "32":
		return "DS2780 - Fuel Gauge IC"
	case "33":
		return "DS1961S,DS2432 - 1Kb Memory with SHA"
	case "34":
		return "DS2703 - SHA Battery Authentication"
	case "35":
		return "DS2755 - Fuel Gauge"
	case "36":
		return "DS2740 - Coulomb Counter"
	case "37":
		return "DS1977 - 32Kb Memory"
	case "3D":
		return "DS2781 - Fueld Gauge IC"
	case "3A":
		return "DS2413 - 2 Channel Switch"
	case "3B":
		return "DS1825,MAX31826,MAX31850 - Temperature Sensor"
	case "41":
		return "DS1923,DS1922E,DS1922L,DS1922T - Hygrochrons"
	case "42":
		return "DS28EA00 - Digital Thermometer With Sequence Detect"
	case "43":
		return "DS28EC20 - 20Kb Memory"
	case "44":
		return "DS28E10 - SHA1 Authenticator"
	case "51":
		return "DS2751 - Battery Fuel Gauge"
	default:
		return fam
	}
}
