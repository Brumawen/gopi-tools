package gopitools

import (
	"os/exec"
	"strconv"
	"strings"
)

type PiStatus struct {
	HostName  string
	IpAddress []string
	CpuTemp   float64
	GpuTemp   float64
}

func (s *PiStatus) Read() error {
	//get hostname
	out, err := exec.Command("hostname").Output()
	if err != nil {
		return err
	}
	s.HostName = string(out)

	//get ip addresses
	ip, err := GetLocalIPAddresses()
	if err != nil {
		return err
	}
	s.IpAddress = ip

	//get cpu temperature
	txt, err := ReadAllText("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return err
	}
	v, err := strconv.ParseFloat(strings.TrimSpace(txt), 64)
	if err == nil {
		s.CpuTemp = v / 1000
	}

	//get gpu temperature
	out, err = exec.Command("/opt/vc/bin/vcgencmd", "measure_temp").Output()
	if err != nil {
		return nil
	}
	txt = string(out)
	txt = txt[5 : len(txt)-3]
	v, err = strconv.ParseFloat(txt, 64)
	if err == nil {
		s.GpuTemp = v
	}
	return nil
}
