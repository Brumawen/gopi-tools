package gopitools

import (
	"fmt"
	"testing"
)

func TestCanGetStatus(t *testing.T) {
	s := PiStatus{}
	s.Read()
	fmt.Println(s)
}
