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
	HostName     string
	IpAddress    []string
	CpuTemp      float64
	GpuTemp      float64
	FreeDisk     int
	FreeDiskPerc int
	AvailMem     int
	Uptime       int
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
		s.IpAddress = ip
	}

	//get cpu temperature
	if txt, err := ReadAllText("/sys/class/thermal/thermal_zone0/temp"); err != nil {
		log.Println("Could not get CPU temperature.", err)
	} else {
		if v, err := strconv.ParseFloat(strings.TrimSpace(txt), 64); err != nil {
			log.Println("Could not parse CPU Temperature.", txt)
		} else {
			s.CpuTemp = v / 1000
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
			s.GpuTemp = v
		}
	}

	//get disk space
	re := regexp.MustCompile("/dev/root\\s*(\\d*)\\s*(\\d*)\\s*(\\d*)\\s*(\\d*)%")
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
