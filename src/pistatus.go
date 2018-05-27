package gopitools

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// PiStatus holds current information about the Device
type PiStatus struct {
	HostName     string   // Current Host Name
	IPAddress    []string // Current IP Address
	OSName       string   // Operating System Name
	OSVersion    string   // Operating System version
	HWType       string   // Hardware type
	HWSerialNo   string   // Hardware SerialNo
	MachineID    string   // Machine ID
	CPUTemp      float64  // CPU temperature in Celcius
	GPUTemp      float64  // GPU temperature in Celcius
	FreeDisk     int      // Free Disk Space in bytes
	FreeDiskPerc int      // Free Disk Space in percentage
	AvailMem     int      // Available Memory in bytes
	Uptime       int      // CPU uptime in seconds

}

// Read loads the status values into the struct
func (s *PiStatus) Read() {
	//get hostname
	if out, err := exec.Command("hostname").Output(); err != nil {
		log.Println("Could not get HostName.", err)
	} else {
		s.HostName = strings.TrimSpace(string(out))
	}

	//get ip addresses
	if ip, err := GetLocalIPAddresses(); err != nil {
		log.Println("Could not get IP addresses.", err)
	} else {
		s.IPAddress = ip
	}

	//get operating system info
	re := regexp.MustCompile("PRETTY_NAME=\"([\\w\\d\\s/()]+)\"")
	if txt, err := ReadAllText("/etc/os-release"); err != nil {
		log.Println("Could not get OS Information")
	} else {
		m := re.FindStringSubmatch(txt)
		if len(m) >= 2 {
			s.OSName = m[1]
		}
	}
	if txt, err := ReadAllText("/etc/debian_version"); err != nil {
		log.Println("Could not get OS Version")
	} else {
		s.OSVersion = strings.TrimSpace(txt)
	}

	//get hardware type
	re = regexp.MustCompile("Revision\\s:\\s([a-e\\d]+)")
	re1 := regexp.MustCompile("Serial\\s\\s:\\s([a-f\\d]*)")
	if txt, err := ReadAllText("/proc/cpuinfo"); err != nil {
		log.Println("Could not get Hardware Type")
	} else {
		m := re.FindStringSubmatch(txt)
		if len(m) >= 2 {
			s.HWType = getHardwareType(m[1])
		}
		m = re1.FindStringSubmatch(txt)
		if len(m) >= 2 {
			s.HWSerialNo = m[1]
		}
	}

	//get machine id
	if txt, err := ReadAllText("/etc/machine-id"); err != nil {
		log.Println("Could not get Machine ID")
	} else {
		s.MachineID = txt
	}

	//get cpu temperature
	if txt, err := ReadAllText("/sys/class/thermal/thermal_zone0/temp"); err != nil {
		log.Println("Could not get CPU temperature.", err)
	} else {
		if v, err := strconv.ParseFloat(strings.TrimSpace(txt), 64); err != nil {
			log.Println("Could not parse CPU Temperature.", txt)
		} else {
			s.CPUTemp = v / 1000
		}
	}

	//get gpu temperature
	if out, err := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp").Output(); err != nil {
		log.Println("Could not get GPU temperature.", err)
	} else {
		txt := string(out)
		txt = txt[5 : len(txt)-3]
		if v, err := strconv.ParseFloat(txt, 64); err != nil {
			log.Println("Could not parse GPU temperature.", txt)
		} else {
			s.GPUTemp = v
		}
	}

	//get disk space
	re = regexp.MustCompile("/dev/root\\s*(\\d*)\\s*(\\d*)\\s*(\\d*)\\s*(\\d*)%")
	if out, err := exec.Command("df").Output(); err != nil {
		log.Println("Could not get Disk Space.", err)
	} else {
		m := re.FindStringSubmatch(string(out))
		if len(m) >= 5 {
			if v, err := strconv.Atoi(m[3]); err != nil {
				log.Println("Could not parse Free Disk Space.", m[3])
			} else {
				s.FreeDisk = v
			}
			if v, err := strconv.Atoi(m[4]); err != nil {
				log.Println("Could not parse Free Disk Percentage.", m[4])
			} else {
				s.FreeDiskPerc = v
			}
		}
	}

	//get available memory
	re = regexp.MustCompile("MemAvailable:\\s*(\\d*)")
	if txt, err := ReadAllText("/proc/meminfo"); err != nil {
		log.Println("Could not get available memory", err)
	} else {
		m := re.FindStringSubmatch(txt)
		if len(m) >= 2 {
			if v, err := strconv.Atoi(m[1]); err != nil {
				log.Println("Could not parse Availabme Memory.", m[2])
			} else {
				s.AvailMem = v
			}
		}
	}

	//get uptime
	if txt, err := ReadAllText("/proc/uptime"); err != nil {
		log.Println("Could not read the system uptime.", err)
	} else {
		i := strings.IndexRune(txt, '.')
		fmt.Println(txt, txt[:i])
		if i >= 0 {
			if v, err := strconv.Atoi(txt[:i]); err != nil {
				log.Println("Could not parse the uptime value.", txt[:i])
			} else {
				s.Uptime = v
			}
		}
	}
}

func getHardwareType(code string) string {
	switch code {
	case "0002", "0003":
		return "Raspberry Pi B rev 1.0"
	case "0004", "0005", "0006", "000d", "000e", "000f":
		return "Raspberry Pi B rev 2.0"
	case "0007", "0008", "0009":
		return "Raspberry Pi A"
	case "0010":
		return "Raspberry Pi B+"
	case "0011":
		return "Raspberry Pi Compute Module"
	case "0012":
		return "Raspberry Pi A+"
	case "a01041", "a21041":
		return "Raspberry Pi 2B"
	case "900092", "900093":
		return "Raspberry Pi Zero"
	case "a02082", "a22082":
		return "Raspberry Pi 3B"
	case "9000c1":
		return "Raspberry Pi Zero W"
	default:
		return "Unknown Model"
	}
}
